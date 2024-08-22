package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/wahyusa/goartisan/internal/generator"
)

var gitFlag bool

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project",
	Run:   runInit,
}

func init() {
	initCmd.Flags().BoolVarP(&gitFlag, "git", "g", false, "Initialize a new git repository")
	baseCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		if err := cmd.Help(); err != nil {
			color.Red("Error displaying help: %v", err)
		}
		return
	}

	projectName := args[0]
	projectPath := filepath.Join(".", projectName)

	// Check if project directory exists
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		color.Red("Project directory already exists")
		os.Exit(1)
	}

	// Initialize spinner
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	// Generate project files
	if err := generator.GenerateProjectFiles(projectName, projectPath, gitFlag); err != nil {
		panic(err)
	}

	// Stop spinner and print success message
	s.Stop()

	// Print next steps
	color.Yellow("🚀 Next steps:")
	color.Magenta("1. Change directory to your project: %s", color.CyanString(fmt.Sprintf("cd %s", projectName)))
	color.Magenta("2. Configure your environment variables in the .env file. 📌")
	color.Magenta("3. To run your project, use %s", color.CyanString("go run main.go"))
	color.Magenta("4. Create models, repositories, services, handlers, etc. 🛠️")
	color.Magenta("5. Add new routes in the routes.go file. 🛤️")
	color.Magenta("6. Re-run your project again. 🔄")
	color.Yellow("💡 Can I use air for hot reload? %s", color.GreenString("Yes!"))

	// Exit
	os.Exit(0)
}
