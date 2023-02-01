package DomainUsers

import (
	DomainUsersEntity "gontact/Domain/Users/Entity"
	DomainUsersVO "gontact/Domain/Users/VO"
	InterfaceUsers "gontact/Interface/Users"
)

func GetUser(vo DomainUsersVO.UserVO) DomainUsersEntity.UsersEntity {
	return getUserEntity(InterfaceUsers.GetDbUser(vo))
}
func GetUserContacts(vo DomainUsersVO.UserContactVO) []DomainUsersEntity.UserContacts {
	return getUserContactsCollection(InterfaceUsers.GetDbUserContacts(vo))
}
