package validations

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations() {
	v := validator.New()

	err := v.RegisterValidation("password", PasswordValidator)
	if err != nil {
		log.Printf("Error registering password validator: %v", err)
	}
}