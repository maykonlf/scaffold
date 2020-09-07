package template

import "github.com/maykonlf/scaffold/internal/domain/entities"

type ConfigService interface {
	GetTemplates() ([]*entities.Template, error)
	AddTemplate(template *entities.Template) error
}
