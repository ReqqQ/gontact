package DomainUsers

import (
	DomainUsersEntity "gontact/Domain/Users/Entity"
	DomainUsersVO "gontact/Domain/Users/VO"
	InterfaceUsers "gontact/Interface/Users"
)

func GetUser(vo DomainUsersVO.UserVO) DomainUsersEntity.UsersEntity {
	return getUserEntity(InterfaceUsers.GetDbUser(vo, false))
}
func GetUserByToken(vo DomainUsersVO.UserTokenVO) DomainUsersEntity.UsersEntity {
	return getUserEntity(InterfaceUsers.GetDbUser(vo, true))
}
func GetUserContacts(vo DomainUsersVO.UserContactVO) []DomainUsersEntity.UsersContacts {
	return getUserContactsCollection(InterfaceUsers.GetDbUserContacts(vo))
}
func CreateUserContact(entity DomainUsersEntity.UsersContacts) {
	InterfaceUsers.CreateUserContact(entity)
}
func IsUserContactExists(entity DomainUsersEntity.UsersContacts, collection []DomainUsersEntity.UsersContacts) bool {
	isExists := false
	for _, dbEntity := range collection {
		if entity.GetEmail() == dbEntity.GetEmail() {
			isExists = true
			break
		}
	}

	return isExists
}
