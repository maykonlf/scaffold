package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainer_GetRootCommand(t *testing.T) {
	container := NewContainer()

	rootCmd := container.GetRootCommand()
	assert.Equal(t, rootCmd, container.GetRootCommand())
}
