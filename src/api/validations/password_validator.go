package validations

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

// PasswordValidator validates the password field according to specific criteria.
// It ensures the password is at least 8 characters long and includes at least one uppercase letter, one lowercase letter, one number, and one special character.
//
// Allowed special characters: !@#$%^&*()_+-=[]{}|;:,.<>?~
//
// Returns true if the password meets all criteria; otherwise, returns false.
func PasswordValidator(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if !ok || len(password) < 8 {
		return false
	}

	hasUpper, hasLower, hasNumber, hasSpecial := false, false, false, false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
		if hasUpper && hasLower && hasNumber && hasSpecial {
			return true
		}
	}

	return false
}
