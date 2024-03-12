package Validators

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

type ValidationError struct {
	errs []error
}

func NewValidationError(errs []error) *ValidationError {
	return &ValidationError{
		errs: errs,
	}
}

func (err *ValidationError) Error() string {
	msg := ""
	for _, err := range err.errs {
		msg += fmt.Sprintf("%s\n", err.Error())
	}
	return msg
}

func (err *ValidationError) Errors() []error {
	return err.errs
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator(validator *validator.Validate) *CustomValidator {
	return &CustomValidator{
		validator: validator,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errs := make([]error, 0, len(validationErrors))
			for _, fieldError := range validationErrors {
				errs = append(errs, fmt.Errorf("%s: violates the %s rule", fieldError.Field(), fieldError.ActualTag()))
			}
			return &ValidationError{
				errs: errs,
			}
		}
	}
	return nil
}
