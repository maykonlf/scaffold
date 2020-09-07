package configs

import "github.com/maykonlf/scaffold/internal/domain/entities"

type Template struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
}

func (t *Template) ToEntity() *entities.Template {
	return &entities.Template{
		Name:   t.Name,
		Source: t.Source,
	}
}

type Config struct {
	Templates []*Template `yaml:"templates"`
}
