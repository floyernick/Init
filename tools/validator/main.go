package validator

import (
	validatorV9 "gopkg.in/go-playground/validator.v9"
)

var validator = validatorV9.New()

func Process(i interface{}) error {
	return validator.Struct(i)
}
