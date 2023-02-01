package AppContact

import (
	AppContactGetGroupTypesCommand "gontact/App/Contact/GetGroupTypes/Command"
	AppContactGetGroupTypesDTO "gontact/App/Contact/GetGroupTypes/DTO"
	AppContactGetUserCommand "gontact/App/Contact/GetUser/Commands"
	AppContactGetUserDTO "gontact/App/Contact/GetUser/DTO"
	AppContactGetUserContactsCommand "gontact/App/Contact/GetUserContacts/Command"
	AppContactGetUserContactsDTO "gontact/App/Contact/GetUserContacts/DTO"
	DomainGroupTypes "gontact/Domain/GroupTypes"
	DomainUsers "gontact/Domain/Users"
)

func GetUser(command AppContactGetUserCommand.UserCommand) AppContactGetUserDTO.UserDTO {
	entity := DomainUsers.GetUser(getUserVO(command))

	return getUserDTO(entity)
}
func GetUserContacts(command AppContactGetUserContactsCommand.UserContactsCommand) []AppContactGetUserContactsDTO.UserContactDTO {
	collection := DomainUsers.GetUserContacts(getUserContactVO(command))

	return getUserContactsDTOCollection(collection)
}
func GetGroupTypes(command AppContactGetGroupTypesCommand.GroupTypesCommand) []AppContactGetGroupTypesDTO.GroupTypesDTO {
	collection := DomainGroupTypes.GetGroupTypes(getGroupTypesVO(command))

	return getGroupTypesDTOCollection(collection)
}
