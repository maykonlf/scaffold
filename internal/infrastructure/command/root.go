package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCommand(templateCommand TemplateCommand) RootCommand {
	rootCmd := &cobra.Command{
		Use:   "scaffold [COMMAND]",
		Short: "Scaffold is template based init tool",
		Long:  "A simple and flexible scaffold tool to initialize projects from templates",
	}

	return &rootCommand{
		cmd:             rootCmd,
		templateCommand: templateCommand,
	}
}

type RootCommand interface {
	Execute()
}

type rootCommand struct {
	cmd             *cobra.Command
	templateCommand TemplateCommand
}

func (c *rootCommand) Execute() {
	c.init()

	if err := c.cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *rootCommand) init() {
	c.cmd.AddCommand(c.templateCommand.List())
	c.cmd.AddCommand(c.templateCommand.Add())
}
