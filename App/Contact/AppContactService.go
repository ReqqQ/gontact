package AppContact

import (
	AppContactGetGroupTypesCommand "gontact/App/Contact/Commands"
	AppContactGetUserContactsDTO "gontact/App/Contact/DTO"
	DomainGroupTypes "gontact/Domain/GroupTypes"
	DomainUsers "gontact/Domain/Users"
)

func GetUser(command AppContactGetGroupTypesCommand.UserCommand) AppContactGetUserContactsDTO.UserDTO {
	entity := DomainUsers.GetUser(getUserVO(command))

	return getUserDTO(entity)
}
func GetUserContacts(command AppContactGetGroupTypesCommand.UserContactsCommand) []AppContactGetUserContactsDTO.UserContactDTO {
	collection := DomainUsers.GetUserContacts(getUserContactVO(command))

	return getUserContactsDTOCollection(collection)
}
func GetGroupTypes(command AppContactGetGroupTypesCommand.GroupTypesCommand) []AppContactGetUserContactsDTO.GroupTypesDTO {
	collection := DomainGroupTypes.GetGroupTypes(getGroupTypesVO(command))

	return getGroupTypesDTOCollection(collection)
}
func CreateUserContact(command AppContactGetGroupTypesCommand.CreateContactCommand) bool {
	entity := GetUserContactEntity(command)
	collection := DomainUsers.GetUserContacts(getUserContactVO(command))
	if DomainUsers.IsUserContactExists(entity, collection) {
		return false
	}

	DomainUsers.CreateUserContact(entity)

	return true
}
