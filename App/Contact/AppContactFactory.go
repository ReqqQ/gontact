package AppContact

import (
	"github.com/gofiber/fiber/v2"
	AppContactGetGroupTypesCommand "gontact/App/Contact/Commands"
	AppContactGetUserContactsDTO "gontact/App/Contact/DTO"
	AppSecuritySecurityDTO "gontact/App/Security/DTO"
	DomainGroupTypesEntity "gontact/Domain/GroupTypes/Entity"
	DomainGroupTypesVO "gontact/Domain/GroupTypes/VO"
	DomainUsers "gontact/Domain/Users"
	DomainUsersEntity "gontact/Domain/Users/Entity"
	DomainUsersVO "gontact/Domain/Users/VO"
)

const (
	ParamNameGroupId = "groupId"
	EmptyString      = ""
)

func GetUserCommand(dto AppSecuritySecurityDTO.TokenUserDTO) AppContactGetGroupTypesCommand.UserCommand {
	return AppContactGetGroupTypesCommand.UserCommand{
		UserId: dto.GetId(),
	}
}

func GetUserContactCommand(c *fiber.Ctx, dto AppSecuritySecurityDTO.TokenUserDTO) AppContactGetGroupTypesCommand.UserContactsCommand {
	var command AppContactGetGroupTypesCommand.UserContactsCommand
	if c.Params(ParamNameGroupId) != EmptyString {
		c.ParamsParser(&command)
	}

	return command.SetUserId(dto.GetId())
}
func GetCreateContactCommand(c *fiber.Ctx, dto AppSecuritySecurityDTO.TokenUserDTO) AppContactGetGroupTypesCommand.CreateContactCommand {
	var command AppContactGetGroupTypesCommand.CreateContactCommand
	err := c.BodyParser(&command)
	if err != nil {
		return AppContactGetGroupTypesCommand.CreateContactCommand{}
	}

	return command.SetUserId(dto.GetId())
}

func GetCreateUser(c *fiber.Ctx) AppContactGetGroupTypesCommand.CreateUserCommand {
	var command AppContactGetGroupTypesCommand.CreateUserCommand
	err := c.BodyParser(&command)
	if err != nil {
		return AppContactGetGroupTypesCommand.CreateUserCommand{}
	}

	return command
}

func GetUserContactEntity(command AppContactGetGroupTypesCommand.CreateContactCommand) DomainUsersEntity.UsersContacts {
	return DomainUsersEntity.UsersContacts{
		UserId:  command.GetUserId(),
		Name:    command.GetUserName(),
		Surname: command.GetUserSurname(),
		Email:   command.GetEmail(),
		Phone:   command.GetUserPhone(),
	}
}
func GetUserEntity(command AppContactGetGroupTypesCommand.CreateUserCommand) DomainUsersEntity.UsersEntity {
	return DomainUsersEntity.UsersEntity{
		Name:     command.GetName(),
		Surname:  command.GetSurname(),
		Email:    command.GetEmail(),
		Password: DomainUsers.HashPassword(command.GetPassword()),
		Token:    DomainUsers.GenerateUserToken(),
	}
}
func GetGroupTypesCommand(dto AppSecuritySecurityDTO.TokenUserDTO) AppContactGetGroupTypesCommand.GroupTypesCommand {
	return AppContactGetGroupTypesCommand.GroupTypesCommand{
		UserId: dto.GetId(),
	}
}

func getUserVO(command AppContactGetGroupTypesCommand.UserCommand) DomainUsersVO.UserVO {
	return DomainUsersVO.UserVO{
		UserId: command.GetUserId(),
	}
}
func getUserContactVO(commandInterface AppContactGetGroupTypesCommand.UserCommandInterface) DomainUsersVO.UserContactVO {
	return DomainUsersVO.UserContactVO{
		UserId:  commandInterface.GetUserId(),
		GroupId: commandInterface.GetGroupId(),
	}
}
func getGroupTypesVO(command AppContactGetGroupTypesCommand.GroupTypesCommand) DomainGroupTypesVO.GroupTypesVO {
	return DomainGroupTypesVO.GroupTypesVO{
		UserId: command.GetUserId(),
	}
}
func getUserDTO(entity DomainUsersEntity.UsersEntity) AppContactGetUserContactsDTO.UserDTO {
	return AppContactGetUserContactsDTO.UserDTO{
		Id:      entity.GetId(),
		Name:    entity.GetName(),
		Surname: entity.GetSurname(),
	}
}
func getUserContactsDTOCollection(entityCollection []DomainUsersEntity.UsersContacts) []AppContactGetUserContactsDTO.
	UserContactDTO {
	collection := []AppContactGetUserContactsDTO.UserContactDTO{}

	for _, entity := range entityCollection {
		collection = append(collection, AppContactGetUserContactsDTO.UserContactDTO{
			Id:      entity.GetId(),
			UserId:  entity.GetUserId(),
			Name:    entity.GetName(),
			Surname: entity.GetSurname(),
			Email:   entity.GetEmail(),
			Phone:   entity.GetPhone(),
		})
	}

	return collection
}
func getGroupTypesDTOCollection(entityCollection []DomainGroupTypesEntity.GroupTypesEntity) []AppContactGetUserContactsDTO.GroupTypesDTO {
	collection := []AppContactGetUserContactsDTO.GroupTypesDTO{}

	for _, entity := range entityCollection {
		collection = append(collection, AppContactGetUserContactsDTO.GroupTypesDTO{
			Id:        entity.GetId(),
			Name:      entity.GetName(),
			CreatedBy: entity.GetCreatedBy(),
		})
	}
	return collection
}
