package routes

import (
	"github.com/gofiber/fiber/v2"
	"gontact/API/groups"
	"gontact/API/users"
)

const (
	apiPrefix             = "/api/"
	endpointUser          = "user/:userId"
	endpointUserContacts  = "user/:userId/contacts"
	endpointContactsGroup = "user/:userId/contacts/group/:groupId"
	endpointGroupType     = "user/:userId/contacts/group/types"
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
	api.Get(endpointGroupType, func(c *fiber.Ctx) error {
		groupTypeVO := groups.GetGroupTypeVO(users.GetUserParamId(c))

		return c.JSON(groups.GetUserGroupTypes(groupTypeVO))
	})
	api.Get(endpointContactsGroup, func(c *fiber.Ctx) error {
		userVo := users.CreateUserVO(users.GetUserParamId(c), users.GetGroupParamId(c))

		return c.JSON(users.GetUserContacts(userVo))
	})
}
func GetPostRoutes(app *fiber.App) {
	api := app.Group(apiPrefix)
	api.Post(endpointUserContacts, func(c *fiber.Ctx) error {
		vo := users.CreateUserContactPostVO(users.GetUserParamId(c), c)
		users.CreateUserContact(vo)
		return c.JSON("")
	})
}
