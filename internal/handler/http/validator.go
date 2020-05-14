package http

import (
	"github.com/qilin/crm-api/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		return errors.NewAPI(errors.ErrValidation, err.Error(), "http.validation")
	}
	return nil
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
