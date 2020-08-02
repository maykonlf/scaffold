package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCommand(templateCommand TemplateCommandI) RootCommandI {
	rootCmd := &cobra.Command{
		Use:   "scaffold [COMMAND]",
		Short: "Scaffold is template based init tool",
		Long:  "A simple and flexible scaffold tool to initialize projects from templates",
	}

	return &RootCommand{
		cmd:             rootCmd,
		templateCommand: templateCommand,
	}
}

type RootCommandI interface {
	Execute()
}

type RootCommand struct {
	cmd             *cobra.Command
	templateCommand TemplateCommandI
}

func (c *RootCommand) Execute() {
	c.init()

	if err := c.cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (c *RootCommand) init() {
	c.cmd.AddCommand(c.templateCommand.List())
	c.cmd.AddCommand(c.templateCommand.Add())
}
