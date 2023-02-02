package AppContact

import (
	"github.com/gofiber/fiber/v2"
	AppContactGetGroupTypesCommand "gontact/App/Contact/GetGroupTypes/Command"
	AppContactGetGroupTypesDTO "gontact/App/Contact/GetGroupTypes/DTO"
	AppContactGetUserCommand "gontact/App/Contact/GetUser/Commands"
	AppContactGetUserDTO "gontact/App/Contact/GetUser/DTO"
	AppContactGetUserContactsCommand "gontact/App/Contact/GetUserContacts/Command"
	AppContactGetUserContactsDTO "gontact/App/Contact/GetUserContacts/DTO"
	AppSecuritySecurityDTO "gontact/App/Security/DTO"
	DomainGroupTypesEntity "gontact/Domain/GroupTypes/Entity"
	DomainGroupTypesVO "gontact/Domain/GroupTypes/VO"
	DomainUsersEntity "gontact/Domain/Users/Entity"
	DomainUsersVO "gontact/Domain/Users/VO"
)

const (
	ParamNameGroupId = "groupId"
	EmptyString      = ""
)

func GetUserCommand(dto AppSecuritySecurityDTO.TokenUserDTO) AppContactGetUserCommand.UserCommand {
	return AppContactGetUserCommand.UserCommand{
		UserId: dto.GetId(),
	}
}

func GetUserContactCommand(c *fiber.Ctx, dto AppSecuritySecurityDTO.TokenUserDTO) AppContactGetUserContactsCommand.UserContactsCommand {
	var command AppContactGetUserContactsCommand.UserContactsCommand
	if c.Params(ParamNameGroupId) != EmptyString {
		c.ParamsParser(&command)
	}

	return command.SetUserId(dto.GetId())
}
func GetGroupTypesCommand(dto AppSecuritySecurityDTO.TokenUserDTO) AppContactGetGroupTypesCommand.GroupTypesCommand {
	return AppContactGetGroupTypesCommand.GroupTypesCommand{
		UserId: dto.GetId(),
	}
}
func getUserVO(command AppContactGetUserCommand.UserCommand) DomainUsersVO.UserVO {
	return DomainUsersVO.UserVO{
		UserId: command.GetUserId(),
	}
}
func getUserContactVO(command AppContactGetUserContactsCommand.UserContactsCommand) DomainUsersVO.UserContactVO {
	return DomainUsersVO.UserContactVO{
		UserId:  command.GetUserId(),
		GroupId: command.GetGroupId(),
	}
}
func getGroupTypesVO(command AppContactGetGroupTypesCommand.GroupTypesCommand) DomainGroupTypesVO.GroupTypesVO {
	return DomainGroupTypesVO.GroupTypesVO{
		UserId: command.GetUserId(),
	}
}
func getUserDTO(entity DomainUsersEntity.UsersEntity) AppContactGetUserDTO.UserDTO {
	return AppContactGetUserDTO.UserDTO{
		Id:      entity.GetId(),
		Name:    entity.GetName(),
		Surname: entity.GetSurname(),
	}
}
func getUserContactsDTOCollection(entityCollection []DomainUsersEntity.UserContacts) []AppContactGetUserContactsDTO.
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
func getGroupTypesDTOCollection(entityCollection []DomainGroupTypesEntity.GroupTypesEntity) []AppContactGetGroupTypesDTO.GroupTypesDTO {
	collection := []AppContactGetGroupTypesDTO.GroupTypesDTO{}

	for _, entity := range entityCollection {
		collection = append(collection, AppContactGetGroupTypesDTO.GroupTypesDTO{
			Id:        entity.GetId(),
			Name:      entity.GetName(),
			CreatedBy: entity.GetCreatedBy(),
		})
	}
	return collection
}
