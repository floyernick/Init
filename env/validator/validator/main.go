package validator

import (
	"Init/env/validator"
	validatorV9 "gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	v *validatorV9.Validate
}

func (validator Validator) Process(i interface{}) error {
	return validator.v.Struct(i)
}

func NewValidator() validator.Validator {
	return Validator{validatorV9.New()}
}
