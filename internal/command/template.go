package command

import (
	"github.com/maykonlf/scaffold/internal/usecase"
	"github.com/spf13/cobra"
)

func NewTemplateCommand(useCase usecase.TemplateUseCaseI) TemplateCommandI {
	return &TemplateCommand{
		useCase: useCase,
	}
}

type TemplateCommandI interface {
	List() *cobra.Command
	Add() *cobra.Command
}

type TemplateCommand struct {
	useCase usecase.TemplateUseCaseI
}

func (c *TemplateCommand) List() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "list templates",
		Example: "scaffold list",
		Run: func(cmd *cobra.Command, args []string) {
			c.useCase.List()
		},
	}
}

func (c *TemplateCommand) Add() *cobra.Command {
	return &cobra.Command{
		Use:     "add",
		Short:   "add scaffold repository",
		Example: "scaffold add <name> <repository>",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.useCase.Add(args[0], args[1])
		},
	}
}
