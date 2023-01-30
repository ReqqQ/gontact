package users

import UserModels "gontact/API/users/models"

func GetUser(vo UserModels.UsersVO) UserModels.Users {
	return getDbUser(vo)
}
func GetUserContacts(vo UserModels.UsersVO) []UserModels.UsersContacts {
	return getDbUserContacts(vo)
}
func CreateUserContact(vo UserModels.UsersContactPostVO) {
	CreateUserContactInDB(CreateUserContactStruct(vo))
}
