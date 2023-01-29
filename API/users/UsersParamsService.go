package users

import (
	"github.com/gofiber/fiber/v2"
	"gontact/API/utils"
	"gontact/API/validator"
)

const (
	paramUserId  = "userId"
	paramGroupId = "groupId"
)

func GetUserParamId(c *fiber.Ctx) int {
	userParam := c.Params(paramUserId)
	validator.CheckIsNumber(userParam)

	return utils.ConvertToInt(userParam)
}

func GetGroupParamId(c *fiber.Ctx) *int {
	userParam := c.Params(paramGroupId)
	validator.CheckIsNumber(userParam)
	convertedParam := utils.ConvertToInt(userParam)

	return &convertedParam
}
