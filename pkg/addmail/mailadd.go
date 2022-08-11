package addmail

import (
	"Logging_Using_ZAP/pkg/logger"
	"Logging_Using_ZAP/pkg/logger2"
	"net/mail"
)

func validMailAddress(address string)  bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}

func PrintValueParse(e string) {
	value := validMailAddress (e)
	if value == true {
		logger.Info("VALID")
		logger2.Info("VALID")
	} else {
		logger.Info("NOT VALID")
		logger2.Error("NOT VALID")
		logger2.Debug("NOT VALID")
		logger2.Warn("NOT VALID")

	}
}