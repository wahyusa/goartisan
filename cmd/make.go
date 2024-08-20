package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wahyusa/goartisan/internal/generator"
)

var makeModelCmd = &cobra.Command{
	Use:   "make:model [name]",
	Short: "Generate a new model",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]
		if err := generator.GenerateModel(modelName); err != nil {
			fmt.Println("Error generating model:", err)
		} else {
			fmt.Println("Model generated successfully:", modelName)
		}
	},
}

func init() {
	baseCmd.AddCommand(makeModelCmd)
}
