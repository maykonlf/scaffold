package configs

import (
	"os"
	"testing"

	"github.com/maykonlf/scaffold/internal/domain/entities"
	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestConfigService(t *testing.T) {
	homeDir, err := homedir.Dir()
	assert.Nil(t, err)

	testConfigDir := homeDir + "/.scaffold-test"
	service := NewService(testConfigDir)

	t.Run("should add template", func(t *testing.T) {
		err := service.AddTemplate(&entities.Template{
			Name:   "template01",
			Source: "source01",
		})

		assert.Nil(t, err)

		err = service.AddTemplate(&entities.Template{
			Name:   "template02",
			Source: "source02",
		})

		assert.Nil(t, err)
	})

	t.Run("should list templates", func(t *testing.T) {
		templates, err := service.GetTemplates()

		assert.Nil(t, err)
		assert.Equal(t, 2, len(templates))
	})

	t.Run("should return error when the config file is corrupted", func(t *testing.T) {
		f, err := os.OpenFile(testConfigDir+"/config", os.O_WRONLY, 0666)
		assert.Nil(t, err)
		defer f.Close()

		_, err = f.WriteString("*&#@$")
		assert.Nil(t, err)

		err = service.AddTemplate(&entities.Template{
			Name:   "template01",
			Source: "source01",
		})
		assert.NotNil(t, err)

		_, err = service.GetTemplates()
		assert.NotNil(t, err)
	})

	_ = os.RemoveAll(testConfigDir)

	t.Run("should return error when cannot access path", func(t *testing.T) {
		err := NewService("/unknown").AddTemplate(&entities.Template{
			Name:   "name",
			Source: "source",
		})

		assert.NotNil(t, err)
	})
}
