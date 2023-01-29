package UserModels

import GroupsModels "gontact/API/groups/models"

type Users struct {
	Id      int
	Name    string
	Surname string
}
type UsersVO struct {
	UserId  int
	GroupId *int
}

type UsersContacts struct {
	Id      int
	UserId  int
	Name    string
	Surname string
	Email   string
	Phone   string
	Group   GroupsModels.GroupType
}

func (uv UsersVO) GetId() int {
	return uv.UserId
}
func (uv UsersVO) GetGroupId() *int {
	return uv.GroupId
}
