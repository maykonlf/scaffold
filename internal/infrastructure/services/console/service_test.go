package console

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_WriteTable(t *testing.T) {
	mockWriter := &MockWriter{}
	service := NewService(mockWriter)

	t.Run("should return error if cannot write data to writer", func(t *testing.T) {
		mockWriter.On("Write", mock.Anything).Return(12, errors.New("cannot write")).Times(1)

		err := service.WriteTable([][]string{{"first", "second", "third"}})

		assert.NotNil(t, err)
	})

	t.Run("should return error if cannot flush writer before writing", func(t *testing.T) {
		mockWriter.On("Write", mock.Anything).Return(12, nil).Times(1)
		mockWriter.On("Flush").Return(errors.New("cannot flush")).Times(1)

		err := service.WriteTable([][]string{{"first", "second", "third"}})

		assert.NotNil(t, err)
	})

	t.Run("should write table to output writer", func(t *testing.T) {
		mockWriter.On("Write", mock.Anything).Return(12, nil).Times(2)
		mockWriter.On("Flush").Return(nil).Times(1)

		err := service.WriteTable([][]string{
			{"first", "second", "third"},
			{"first-2", "second-2", "third-2"},
		})

		assert.Nil(t, err)
	})
}
