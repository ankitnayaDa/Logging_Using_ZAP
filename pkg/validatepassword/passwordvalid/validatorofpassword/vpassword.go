package validatorofpassword

import (
	"Logging_Using_ZAP/pkg/logger"
	"unicode"
)


func isValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
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
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}


func PrintPassword(e string) {
	value := isValid (e)
	if value == true {
		logger.Info("VALID")
	} else {
		logger.Info("NOT VALID")
	}
}