package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wahyusa/goartisan/internal/generator"
)

func createGenerateCommand(use, short string, generateFunc func(string) error) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			if err := generateFunc(name); err != nil {
				fmt.Println("Error generating", use+":", err)
			} else {
				fmt.Println(use, "generated successfully:", name)
			}
		},
	}
}

var makeModelCmd = createGenerateCommand("make:model", "Generate a new model", generator.GenerateModel)
var makeRepoCmd = createGenerateCommand("make:repo", "Generate a new repository", generator.GenerateRepository)
var makeServiceCmd = createGenerateCommand("make:service", "Generate a new service", generator.GenerateService)
var makeHandlerCmd = createGenerateCommand("make:handler", "Generate a new handler", generator.GenerateHandler)

func init() {
	baseCmd.AddCommand(makeModelCmd)
	baseCmd.AddCommand(makeRepoCmd)
	baseCmd.AddCommand(makeServiceCmd)
	baseCmd.AddCommand(makeHandlerCmd)
}
