package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRequest_Validate(t *testing.T) {
	t.Run("should return error when name or source is empty", func(t *testing.T) {
		r := AddRequest{}
		assert.NotNil(t, r.Validate())

		r = AddRequest{Name: "my-template-repo"}
		assert.NotNil(t, r.Validate())

		r = AddRequest{Name: "", Source: "https://github.com/template/template.git"}
		assert.NotNil(t, r.Validate())
	})

	t.Run("should return nil when the template values is valid", func(t *testing.T) {
		r := AddRequest{Name: "my-template-repo", Source: "https://github.com/template/template.git"}
		assert.Nil(t, r.Validate())

		r = AddRequest{Name: "my-template-repo", Source: "git@github.com:template/template.git"}
		assert.Nil(t, r.Validate())
	})
}
