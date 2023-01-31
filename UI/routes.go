package UI

import (
	"github.com/gofiber/fiber/v2"
	AppContact "gontact/App/Contact"
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
		command := AppContact.GetUserCommand(c).Validate()
		dto := AppContact.GetUser(command)

		return c.JSON(dto)
	})
	api.Get(endpointUserContacts, func(c *fiber.Ctx) error {
		command := AppContact.GetUserContactCommand(c).Validate()
		dtoCollection := AppContact.GetUserContacts(command)

		return c.JSON(dtoCollection)
	})
	api.Get(endpointGroupType, func(c *fiber.Ctx) error {
		command := AppContact.GetGroupTypesCommand(c).Validate()
		dtoCollection := AppContact.GetGroupTypes(command)

		return c.JSON(dtoCollection)
	})
	api.Get(endpointContactsGroup, func(c *fiber.Ctx) error {
		command := AppContact.GetUserContactCommand(c).Validate()
		dtoCollection := AppContact.GetUserContacts(command)

		return c.JSON(dtoCollection)
	})
}
func GetPostRoutes(app *fiber.App) {
	//api := app.Group(apiPrefix)
	//api.Post(endpointUserContacts, func(c *fiber.Ctx) error {
	//	vo := users.CreateUserContactPostVO(users.GetUserParamId(c), c)
	//	users.CreateUserContact(vo)
	//	var i interface{}
	//	return c.JSON(i)
	//})
}
