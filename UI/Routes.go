package UI

import (
	"github.com/gofiber/fiber/v2"
	AppContact "gontact/App/Contact"
	"gontact/App/Response"
	AppSecurity "gontact/App/Security"
	AppSecuritySecurityDTO "gontact/App/Security/DTO"
)

const (
	apiPrefix             = "/api/"
	endpointUser          = "user"
	endpointUserContacts  = "user/contacts"
	endpointContactsGroup = "user/contacts/group/:groupId?"
	endpointGroupType     = "user/contacts/group/types"
)

var userTokenDTO AppSecuritySecurityDTO.TokenUserDTO
var api fiber.Router

func GetRoutes(app *fiber.App) {
	api = app.Group(apiPrefix, middleware)
	api.Get(endpointUser, func(c *fiber.Ctx) error {
		command := AppContact.GetUserCommand(userTokenDTO)
		errors := AppSecurity.Validate(command)
		if errors != nil {
			return c.JSON(Response.Error(errors))
		}
		dto := AppContact.GetUser(command)

		return c.JSON(Response.Success(dto))
	})
	api.Get(endpointUserContacts, func(c *fiber.Ctx) error {
		command := AppContact.GetUserContactCommand(c, userTokenDTO)
		errors := AppSecurity.Validate(command)
		if errors != nil {
			return c.JSON(Response.Error(errors))
		}
		dtoCollection := AppContact.GetUserContacts(command)

		return c.JSON(Response.Success(dtoCollection))
	})
	api.Get(endpointGroupType, func(c *fiber.Ctx) error {
		command := AppContact.GetGroupTypesCommand(userTokenDTO)
		errors := AppSecurity.Validate(command)
		if errors != nil {
			return c.JSON(Response.Error(errors))
		}
		dtoCollection := AppContact.GetGroupTypes(command)

		return c.JSON(Response.Success(dtoCollection))
	})
	api.Get(endpointContactsGroup, func(c *fiber.Ctx) error {
		command := AppContact.GetUserContactCommand(c, userTokenDTO)
		errors := AppSecurity.Validate(command)
		if errors != nil {
			return c.JSON(Response.Error(errors))
		}
		dtoCollection := AppContact.GetUserContacts(command)

		return c.JSON(Response.Success(dtoCollection))
	})
}
func GetPostRoutes() {
	api.Post(endpointUserContacts, func(c *fiber.Ctx) error {
		command := AppContact.GetCreateContactCommand(c, userTokenDTO)
		errors := AppSecurity.Validate(command)
		if errors != nil {
			return c.JSON(Response.Error(errors))
		}
		isCreated := AppContact.CreateUserContact(command)
		if !isCreated {
			return c.JSON(Response.Error(AppContact.ErrorUserAlreadyExists()))
		}
		var i interface{}
		return c.JSON(i)
	})
}
