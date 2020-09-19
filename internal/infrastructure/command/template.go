package command

import (
	"github.com/maykonlf/scaffold/internal/domain/entities"
	"github.com/maykonlf/scaffold/internal/infrastructure/services/console"
	"github.com/maykonlf/scaffold/internal/interface/controllers"
	"github.com/spf13/cobra"
)

func NewTemplateCommand(console console.Service, controller controllers.TemplateController) TemplateCommand {
	return &templateCommand{
		consoleService: console,
		controller:     controller,
	}
}

type TemplateCommand interface {
	List() *cobra.Command
	Add() *cobra.Command
}

type templateCommand struct {
	consoleService console.Service
	controller     controllers.TemplateController
}

func (c *templateCommand) List() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "list templates",
		Example: "scaffold list",
		RunE:    c.listTemplates,
	}
}

func (c *templateCommand) Add() *cobra.Command {
	return &cobra.Command{
		Use:     "add",
		Short:   "add scaffold repository",
		Example: "scaffold add <name> <repository>",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.controller.Add(args[0], args[1])
		},
	}
}

func (c *templateCommand) listTemplates(cmd *cobra.Command, args []string) error {
	templates, err := c.controller.List()
	if err != nil {
		return err
	}

	return c.consoleService.WriteTable(c.parseTemplatesToRows(templates))
}

func (c *templateCommand) parseTemplatesToRows(templates []*entities.Template) [][]string {
	if len(templates) == 0 {
		return [][]string{}
	}

	rows := make([][]string, len(templates)+1)
	rows[0] = []string{"name", "source"}
	for i := range templates {
		rows[i+1] = []string{
			templates[i].Name,
			templates[i].Source,
		}
	}

	return rows
}
