package AppContact

import (
	"github.com/gofiber/fiber/v2"
	AppContactGetGroupTypesCommand "gontact/App/Contact/GetGroupTypes/Command"
	AppContactGetGroupTypesDTO "gontact/App/Contact/GetGroupTypes/DTO"
	AppContactGetUserCommand "gontact/App/Contact/GetUser/Commands"
	AppContactGetUserDTO "gontact/App/Contact/GetUser/DTO"
	AppContactGetUserContactsCommand "gontact/App/Contact/GetUserContacts/Command"
	AppContactGetUserContactsDTO "gontact/App/Contact/GetUserContacts/DTO"
	DomainGroupTypesEntity "gontact/Domain/GroupTypes/Entity"
	DomainGroupTypesVO "gontact/Domain/GroupTypes/VO"
	DomainUsersEntity "gontact/Domain/Users/Entity"
	DomainUsersVO "gontact/Domain/Users/VO"
)

func GetUserCommand(c *fiber.Ctx) AppContactGetUserCommand.UserCommand {
	var command AppContactGetUserCommand.UserCommand
	c.ParamsParser(&command)

	return command
}

func GetUserContactCommand(c *fiber.Ctx) AppContactGetUserContactsCommand.UserContactsCommand {
	var command AppContactGetUserContactsCommand.UserContactsCommand
	c.ParamsParser(&command)

	return command
}
func GetGroupTypesCommand(c *fiber.Ctx) AppContactGetGroupTypesCommand.GroupTypesCommand {
	var command AppContactGetGroupTypesCommand.GroupTypesCommand
	c.ParamsParser(&command)

	return command
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
	var collection []AppContactGetUserContactsDTO.UserContactDTO

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
	var collection []AppContactGetGroupTypesDTO.GroupTypesDTO

	for _, entity := range entityCollection {
		collection = append(collection, AppContactGetGroupTypesDTO.GroupTypesDTO{
			Id:        entity.GetId(),
			Name:      entity.GetName(),
			CreatedBy: entity.GetCreatedBy(),
		})
	}
	return collection
}
