package routes

import (
	"github.com/gofiber/fiber/v2"
	"gontact/API/users"
)

const (
	apiPrefix             = "/api/"
	endpointUser          = "user/:userId"
	endpointUserContacts  = "user/:userId/contacts"
	endpointContactsGroup = "user/:userId/contacts/group/:groupId"
)

func GetRoutes(app *fiber.App) {
	api := app.Group(apiPrefix)
	api.Get(endpointUser, func(c *fiber.Ctx) error {
		userVo := users.CreateUserVO(users.GetUserParamId(c), nil)

		return c.JSON(users.GetUser(userVo))
	})
	api.Get(endpointUserContacts, func(c *fiber.Ctx) error {
		userVo := users.CreateUserVO(users.GetUserParamId(c), nil)

		return c.JSON(users.GetUserContacts(userVo))
	})
	api.Get(endpointContactsGroup, func(c *fiber.Ctx) error {
		userVo := users.CreateUserVO(users.GetUserParamId(c), users.GetGroupParamId(c))

		return c.JSON(users.GetUserContacts(userVo))
	})
}
