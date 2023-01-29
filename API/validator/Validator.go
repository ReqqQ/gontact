package validator

import (
	"log"
	"regexp"
)

const (
	regexOnlyNumbers      = `\d`
	errorParamIsNotNumber = "PARAM_USER_ID_IS_NOT_A_NUMBER"
	errorDefault          = "Error: "
)

func CheckIsNumber(paramValue string) {
	if !regexp.MustCompile(regexOnlyNumbers).MatchString(paramValue) {
		log.Fatal(errorParamIsNotNumber)
	}
}
func DefaultErrorCheck(err error) {
	if err != nil {
		log.Fatal(errorDefault + err.Error())
	}
}
