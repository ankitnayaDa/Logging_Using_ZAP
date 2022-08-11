package main

import (
	"Logging_Using_ZAP/pkg/addmail"
	"Logging_Using_ZAP/pkg/validate/email"
	"Logging_Using_ZAP/pkg/validatepassword/passwordvalid/validatorofpassword"
)

func main(){

	email.PrintValueRegex("email@email.com")
	email.PrintValueRegex("email")

	addmail.PrintValueParse("email@email.com")
	addmail.PrintValueParse("email")
	validatorofpassword.PrintPassword("yt_xk39b")
}
