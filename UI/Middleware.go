package UI

import (
	"github.com/gofiber/fiber/v2"
	"gontact/App/Response"
	AppSecurity "gontact/App/Security"
)

func middleware(c *fiber.Ctx) error {
	securityCommand := AppSecurity.GetSecurityCommand(c)
	errors := AppSecurity.Validate(securityCommand)
	if errors != nil {
		return c.JSON(Response.Error(errors))
	}
	userTokenDTO = AppSecurity.GetCurrentUser(securityCommand)
	if AppSecurity.Validate(userTokenDTO) != nil {
		return c.JSON(Response.Error([]string{Response.UserNotExists}))
	}

	return c.Next()
}
