package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wahyusa/goartisan/internal/config"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project",
	Run:   runInit,
}

func init() {
	baseCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		cmd.Help()
		return
	}

	projectName := args[0]
	projectPath := filepath.Join(".", projectName)

	// Check if project directory exists
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		color.Red("Project directory already exists")
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Initialize spinner
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	// Create directories
	for _, dir := range cfg.Structure.Dirs {
		dirPath := filepath.Join(projectPath, dir)
		err := os.MkdirAll(dirPath, fs.ModeDir)
		if err != nil {
			panic(err)
		}
	}

	// Create files
	for _, filePath := range cfg.Structure.Files {
		fullPath := filepath.Join(projectPath, filePath)
		file, err := os.Create(fullPath)
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	// Change directory to the project path
	if err := os.Chdir(projectPath); err != nil {
		panic(err)
	}

	// Stop spinner and print success message
	s.Stop()
	color.Green("Project initialized successfully")

	// Print project structure
	color.Cyan("Project structure:")
	color.Cyan(projectName)
	for _, dir := range cfg.Structure.Dirs {
		color.Cyan("├── " + dir)
	}
	for _, file := range cfg.Structure.Files {
		color.Cyan("├── " + file)
	}

	// Auto run 'go mod init' command
	color.Yellow("Running 'go mod init' command...")
	modCmd := exec.Command("go", "mod", "init", projectName)
	modCmd.Stdout = os.Stdout
	modCmd.Stderr = os.Stderr
	if err := modCmd.Run(); err != nil {
		panic(err)
	}

	// Print next steps
	color.White("Next steps:")
	color.White("1. Change directory to your project %s", color.BlueString(fmt.Sprintf("cd "+projectName)))
	color.White("2. Start coding your project")
	color.White("3. Adjust .env and database")
	color.White("4. Make model, repository, service, handler etc %s", color.RedString("(WIP)"))
	color.White("3. To start your project run %s", color.BlueString("go run main.go"))

	// Exit
	os.Exit(0)
}
