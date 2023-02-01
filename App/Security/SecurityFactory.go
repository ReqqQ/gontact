package AppSecurity

import (
	"github.com/gofiber/fiber/v2"
	SecurityCommand "gontact/App/Security/Command"
	AppSecuritySecurityDTO "gontact/App/Security/DTO"
	DomainUsersEntity "gontact/Domain/Users/Entity"
)

func GetSecurityCommand(c *fiber.Ctx) SecurityCommand.SecurityCommand {
	var command SecurityCommand.SecurityCommand
	c.ReqHeaderParser(&command)

	return command
}
func GetTokenUserDTO(entity DomainUsersEntity.UsersEntity) AppSecuritySecurityDTO.TokenUserDTO {
	return AppSecuritySecurityDTO.TokenUserDTO{
		Id:      entity.GetId(),
		Name:    entity.GetName(),
		Surname: entity.GetSurname(),
	}
}
