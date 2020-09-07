package template

import (
	"github.com/maykonlf/scaffold/internal/domain/entities"
)

func NewTemplaceUseCase(configService ConfigService) UseCase {
	return &useCase{
		configService: configService,
	}
}

type UseCase interface {
	List() ([]*entities.Template, error)
	Add(request *AddRequest) error
}

type useCase struct {
	configService ConfigService
}

func (u *useCase) List() ([]*entities.Template, error) {
	return u.configService.GetTemplates()
}

func (u *useCase) Add(request *AddRequest) error {
	return u.configService.AddTemplate(&entities.Template{
		Name:   request.Name,
		Source: request.Source,
	})
}
