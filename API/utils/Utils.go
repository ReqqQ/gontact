package utils

import (
	"strconv"

	"gontact/API/validator"
)

func ConvertToInt(paramValue string) int {
	userId, err := strconv.Atoi(paramValue)
	validator.DefaultErrorCheck(err)

	return userId
}
