package cmd

import (
	"github.com/spf13/cobra"
	// "github.com/wahyusa/goartisan/internal/generator"
)

var makeCmd = &cobra.Command{
	// Define the make command
}

func init() {
	// Add make command and subcommands to root command
}

func runMakeAll(cmd *cobra.Command, args []string) {
	// Implement make:all command logic
}

// Add other make subcommands (model, repo, service, handler, middleware)