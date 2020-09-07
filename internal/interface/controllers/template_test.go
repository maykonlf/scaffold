package controllers

import (
	"errors"
	"testing"

	"github.com/maykonlf/scaffold/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTemplatesController_List(t *testing.T) {
	useCase := &UseCase{}
	controller := NewTemplateController(useCase)

	t.Run("should return error when usecase returns an error", func(t *testing.T) {
		mockError := errors.New("mock error")
		useCase.On("List").Return(nil, mockError).Times(1)

		response, err := controller.List()
		assert.Equal(t, mockError, err)
		assert.Nil(t, response)
	})

	t.Run("should return nil when usecase doesn't return error", func(t *testing.T) {
		mockResponse := []*entities.Template{
			{Name: "template-1", Source: "https://source-1"},
			{Name: "template-2", Source: "https://source-2"},
			{Name: "template-3", Source: "https://source-3"},
		}
		useCase.On("List").Return(mockResponse, nil).Times(1)

		response, err := controller.List()
		assert.Nil(t, err)
		assert.Equal(t, mockResponse, response)
	})
}

func TestTemplatesController_Add(t *testing.T) {
	useCase := &UseCase{}
	controller := NewTemplateController(useCase)

	t.Run("should return error when usecase returns an error", func(t *testing.T) {
		mockError := errors.New("mock error")
		useCase.On("Add", mock.Anything).Return(mockError).Times(1)

		err := controller.Add("template-1", "source-1")
		assert.Equal(t, mockError, err)
	})

	t.Run("should return error when usecase doesn't return error", func(t *testing.T) {
		useCase.On("Add", mock.Anything).Return(nil).Times(1)

		err := controller.Add("template-2", "source-2")
		assert.Nil(t, err)
	})
}
