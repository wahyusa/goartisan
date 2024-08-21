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

func GenerateProjectFiles(projectName, projectPath string, gitFlag bool) error {
	// Convert projectPath to an absolute path
	absProjectPath, err := filepath.Abs(projectPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Create directories
	for _, dir := range cfg.Structure.Dirs {
		dirPath := filepath.Join(absProjectPath, dir)
		err := os.MkdirAll(dirPath, fs.ModeDir)
		if err != nil {
			return err
		}
	}

	// Create files
	for _, filePath := range cfg.Structure.Files {
		fullPath := filepath.Join(absProjectPath, filePath)
		file, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		file.Close()
	}

	// Render and write templates
	templates := map[string][]byte{
		"main.go":                  mainTemplate,
		"app/database/database.go": databaseTemplate,
		"app/routes/routes.go":     routesTemplate,
		".env":                     envTemplate,
		".gitignore":               gitignoreTemplate,
	}

	data := map[string]string{
		"ProjectName": projectName,
		"DBDriver":    "mysql", // You can adjust this based on your configuration
		"Port":        "8080",
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

	// Change directory to the project path
	if err := os.Chdir(absProjectPath); err != nil {
		return err
	}

	// Create or open the log file
	logFile := filepath.Join(absProjectPath, "goartisan.log")
	log, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer log.Close()

	// Auto run 'go mod init' command
	color.Yellow("Running 'go mod init' command...")
	modCmd := exec.Command("go", "mod", "init", projectName)
	modCmd.Stdout = log
	modCmd.Stderr = log
	if err := modCmd.Run(); err != nil {
		return err
	}

	// Auto run 'go mod tidy' command
	color.Yellow("Running 'go mod tidy' command...")
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = absProjectPath
	tidyOutput, err := tidyCmd.CombinedOutput()
	if err != nil {
		return err
	}

	// Log the output to the log file
	if _, err := log.Write(tidyOutput); err != nil {
		color.Red("Failed to write go mod tidy log: %v", err)
	}

	// Initialize git repository
	if gitFlag {
		color.Yellow("Initializing git repository...")
		gitCmd := exec.Command("git", "init")
		gitCmd.Dir = absProjectPath
		gitCmd.Stdout = log
		gitCmd.Stderr = log
		if err := gitCmd.Run(); err != nil {
			return err
		}
	}

	// Log project structure
	log.WriteString("Project structure:\n")
	log.WriteString(projectName + "\n")
	for _, dir := range cfg.Structure.Dirs {
		log.WriteString("├── " + dir + "\n")
	}
	for _, file := range cfg.Structure.Files {
		log.WriteString("├── " + file + "\n")
	}

	color.Green("Project initialized successfully.")
	return nil
}
