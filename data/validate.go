package data

import "gopkg.in/go-playground/validator.v9"

func (a *ArtPiece) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("description", validateDesc)
	return validate.Struct(a)
}

func validateDesc(fl validator.FieldLevel) bool {
	descStr := fl.Field().String()

	if len(descStr) > 5 {
		return true
	}

	return false
}
