package data

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// ValidationError wraps the validators FieldError so we do not
// expose this to out code
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("description", validateDesc)

	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {

	error := v.validate.Struct(i)

	if error != nil {
		errs := error.(validator.ValidationErrors)

		if len(errs) == 0 {
			return nil
		}

		var returnErrs []ValidationError

		for _, err := range errs {
			// cast the FieldError into our ValidationError and append to the slice
			ve := ValidationError{err.(validator.FieldError)}
			returnErrs = append(returnErrs, ve)
		}

		return returnErrs
	}
	return nil
}

func validateDesc(fl validator.FieldLevel) bool {
	descStr := fl.Field().String()

	if len(descStr) > 5 {
		return true
	}

	return false
}
