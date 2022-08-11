package email

import (
	"Logging_Using_ZAP/pkg/logger"
	"regexp"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func PrintValueRegex(e string) {
	value := IsEmailValid (e)
	if value == true {
		logger.Info("VALID")
	} else {
		logger.Info("NOT VALID")
	}
}