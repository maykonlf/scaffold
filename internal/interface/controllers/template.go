package controllers

import (
	"github.com/maykonlf/scaffold/internal/domain/entities"
	"github.com/maykonlf/scaffold/internal/usecase/template"
)

type TemplateController interface {
	List() ([]*entities.Template, error)
	Add(name, url string) error
}

func NewTemplateController(useCase template.UseCase) TemplateController {
	return &templateController{
		useCase: useCase,
	}
}

type templateController struct {
	useCase template.UseCase
}

func (c *templateController) List() ([]*entities.Template, error) {
	return c.useCase.List()
}

func (c *templateController) Add(name, source string) error {
	return c.useCase.Add(&template.AddRequest{
		Name:   name,
		Source: source,
	})
}
