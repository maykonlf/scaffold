package template

import (
	"errors"
	"testing"

	"github.com/maykonlf/scaffold/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCase_ListTempates(t *testing.T) {
	configService := &MockConfigService{}
	useCase := NewTemplaceUseCase(configService)

	t.Run("should return error if config service return an error", func(t *testing.T) {
		testError := errors.New("mocked error")
		configService.On("GetTemplates").Return(nil, testError).Times(1)

		template, err := useCase.List()
		assert.Nil(t, template)
		assert.Equal(t, err, testError)
	})

	t.Run("should return templates on success", func(t *testing.T) {
		templates := []*entities.Template{
			{Name: "template-1", Source: "https://templates.com/template1.git"},
			{Name: "template-2", Source: "https://templates.com/template2.git"},
			{Name: "template-3", Source: "https://templates.com/template3.git"},
		}

		configService.On("GetTemplates").Return(templates, nil).Times(1)

		response, err := useCase.List()
		assert.Nil(t, err)
		assert.Equal(t, templates, response)
	})
}

func TestUseCase_Add(t *testing.T) {
	configService := &MockConfigService{}
	useCase := NewTemplaceUseCase(configService)

	t.Run("should return error when fails to add a new template", func(t *testing.T) {
		request := &AddRequest{Name: "template-1", Source: "https://assdasda.com/ddasdsa/33212"}
		mockError := errors.New("mock error")
		configService.On("AddTemplate", mock.Anything).Return(mockError).Times(1)

		err := useCase.Add(request)
		assert.NotNil(t, err)
	})

	t.Run("should return nil when successfully add new template", func(t *testing.T) {
		request := &AddRequest{Name: "template-1", Source: "https://assdasda.com/ddasdsa/33212"}
		configService.On("AddTemplate", mock.Anything).Return(nil).Times(1)

		err := useCase.Add(request)
		assert.Nil(t, err)
	})
}
