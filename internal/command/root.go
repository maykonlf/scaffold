package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scaffold [COMMAND]",
	Short: "Scaffold is template based init tool",
	Long:  "A simple and flexible scaffold tool to initalize projects from templates",
}

func Initialize() {
	rootCmd.AddCommand(templateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
