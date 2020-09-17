package config

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/maykonlf/scaffold/internal/infrastructure/command"
	"github.com/maykonlf/scaffold/internal/infrastructure/services/configs"
	"github.com/maykonlf/scaffold/internal/infrastructure/services/console"
	"github.com/maykonlf/scaffold/internal/interface/controllers"
	"github.com/maykonlf/scaffold/internal/usecase/template"
	"github.com/mitchellh/go-homedir"
)

type Container interface {
	GetRootCommand() command.RootCommand
}

func NewContainer() Container {
	return &container{}
}

type container struct {
	// Infrastructure
	rootCommand     command.RootCommand
	templateCommand command.TemplateCommand
	consoleService  console.Service
	configsService  configs.Service

	// Interface
	templateController controllers.TemplateController

	// Use Case
	templateUseCase template.UseCase
}

func (c *container) GetRootCommand() command.RootCommand {
	if c.rootCommand == nil {
		c.rootCommand = command.NewRootCommand(c.getTemplateCommand())
	}

	return c.rootCommand
}

func (c *container) getTemplateCommand() command.TemplateCommand {
	if c.templateCommand == nil {
		c.templateCommand = command.NewTemplateCommand(c.getConsoleService(), c.getTemplateController())
	}

	return c.templateCommand
}

func (c *container) getConsoleService() console.Service {
	if c.consoleService == nil {
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		c.consoleService = console.NewService(writer)
	}

	return c.consoleService
}

func (c *container) getTemplateController() controllers.TemplateController {
	if c.templateController == nil {
		c.templateController = controllers.NewTemplateController(c.getTemplateUseCase())
	}

	return c.templateController
}

func (c *container) getTemplateUseCase() template.UseCase {
	if c.templateUseCase == nil {
		c.templateUseCase = template.NewTemplaceUseCase(c.getConfigService())
	}

	return c.templateUseCase
}

func (c *container) getConfigService() configs.Service {
	if c.configsService == nil {
		dir, err := homedir.Dir()
		if err != nil {
			panic(err)
		}

		c.configsService = configs.NewService(fmt.Sprintf("%s/.scaffold", dir))
	}

	return c.configsService
}
