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

var makeRepoCmd = &cobra.Command{
	Use:   "make:repo [name]",
	Short: "Generate a new repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoName := args[0]
		if err := generator.GenerateRepository(repoName); err != nil {
			fmt.Println("Error generating repository:", err)
		} else {
			fmt.Println("Repository generated successfully:", repoName)
		}
	},
}

var makeServiceCmd = &cobra.Command{
	Use:   "make:service [name]",
	Short: "Generate a new service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		if err := generator.GenerateService(serviceName); err != nil {
			fmt.Println("Error generating service:", err)
		} else {
			fmt.Println("Service generated successfully:", serviceName)
		}
	},
}

var makeHandlerCmd = &cobra.Command{
	Use:   "make:handler [name]",
	Short: "Generate a new handler",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handlerName := args[0]
		if err := generator.GenerateHandler(handlerName); err != nil {
			fmt.Println("Error generating handler:", err)
		} else {
			fmt.Println("Handler generated successfully:", handlerName)
		}
	},
}

func init() {
	baseCmd.AddCommand(makeModelCmd)
	baseCmd.AddCommand(makeRepoCmd)
	baseCmd.AddCommand(makeServiceCmd)
	baseCmd.AddCommand(makeHandlerCmd)
}
