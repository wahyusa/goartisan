package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	"github.com/wahyusa/goartisan/internal/config"
)

// Embed the templates
//
//go:embed templates/main.go.tmpl
var mainTemplate []byte

//go:embed templates/database.go.tmpl
var databaseTemplate []byte

//go:embed templates/routes.go.tmpl
var routesTemplate []byte

//go:embed templates/env.go.tmpl
var envTemplate []byte

//go:embed templates/gitignore.go.tmpl
var gitignoreTemplate []byte

//go:embed templates/default_config.toml.tmpl
var defaultConfig []byte

func GenerateProjectFiles(projectName, projectPath string, gitFlag bool) error {
	absProjectPath, err := filepath.Abs(projectPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	if err := ensureConfigDirectory(); err != nil {
		return err
	}

	if err := loadAndSaveConfig(projectName); err != nil {
		return err
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	if err := createProjectStructure(absProjectPath, cfg); err != nil {
		return err
	}

	if err := renderTemplates(absProjectPath, projectName, cfg); err != nil {
		return err
	}

	if err := initializeGoModules(absProjectPath, projectName); err != nil {
		return err
	}

	if gitFlag {
		if err := initializeGitRepository(absProjectPath); err != nil {
			return err
		}
	}

	color.Green("Project initialized successfully.")
	return nil
}

func ensureConfigDirectory() error {
	configDir := filepath.Join(os.Getenv("HOME"), ".goartisan")
	if err := os.MkdirAll(configDir, fs.ModePerm); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	configFilePath := filepath.Join(configDir, "config.toml")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		if err := os.WriteFile(configFilePath, defaultConfig, 0644); err != nil {
			return fmt.Errorf("failed to write default config file: %v", err)
		}
	}
	return nil
}

func loadAndSaveConfig(projectName string) error {
	configFilePath := filepath.Join(os.Getenv("HOME"), ".goartisan", "config.toml")
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	viper.Set("module.name", projectName)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}
	return nil
}

func createProjectStructure(absProjectPath string, cfg *config.Config) error {
	for _, dir := range cfg.Structure.Dirs {
		dirPath := filepath.Join(absProjectPath, dir)
		if err := os.MkdirAll(dirPath, fs.ModeDir); err != nil {
			return err
		}
	}

	for _, filePath := range cfg.Structure.Files {
		fullPath := filepath.Join(absProjectPath, filePath)
		file, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func renderTemplates(absProjectPath, projectName string, cfg *config.Config) error {
	templates := map[string][]byte{
		"main.go":                  mainTemplate,
		"app/database/database.go": databaseTemplate,
		"app/routes/routes.go":     routesTemplate,
		".env":                     envTemplate,
		".gitignore":               gitignoreTemplate,
	}

	data := map[string]string{
		"ProjectName": projectName,
		"DBDriver":    cfg.Database.Default,
		"Port":        cfg.Server.Port,
	}

	for filePath, tmplContent := range templates {
		fullPath := filepath.Join(absProjectPath, filePath)
		tmpl, err := template.New(filePath).Parse(string(tmplContent))
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return err
		}

		if err := os.WriteFile(fullPath, buf.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

func initializeGoModules(absProjectPath, projectName string) error {
	logFile := filepath.Join(absProjectPath, "goartisan.log")
	log, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer log.Close()

	if err := runCommand("go mod init", absProjectPath, log, "go", "mod", "init", projectName); err != nil {
		return err
	}

	if err := runCommand("go mod tidy", absProjectPath, log, "go", "mod", "tidy"); err != nil {
		return err
	}

	return nil
}

func initializeGitRepository(absProjectPath string) error {
	logFile := filepath.Join(absProjectPath, "goartisan.log")
	log, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer log.Close()

	return runCommand("git init", absProjectPath, log, "git", "init")
}

func runCommand(description, dir string, log *os.File, name string, args ...string) error {
	color.Yellow("Running '%s' command...", description)
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = log
	cmd.Stderr = log
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run %s: %v", description, err)
	}
	return nil
}
