package template

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type AddRequest struct {
	Name   string
	Source string
}

func (r AddRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Source, validation.Required, validation.Match(regexp.MustCompile(`^(https://|git@).{6,}(\\.git)?`))),
	)
}
