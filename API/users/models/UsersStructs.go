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

type UsersContactPostVO struct {
	UserId  int
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type UsersContacts struct {
	Id      int
	UserId  int
	Name    string
	Surname string
	Email   string
	Phone   string
	Group   GroupsModels.GroupType `gorm:"-"`
}

func (uv UsersVO) GetId() int {
	return uv.UserId
}
func (uv UsersVO) GetGroupId() *int {
	return uv.GroupId
}
func (r UsersContactPostVO) SetUserId(userId int) UsersContactPostVO {
	r.UserId = userId

	return r
}
func (r UsersContactPostVO) GetUserId() int {
	return r.UserId
}
func (r UsersContactPostVO) GetName() string {
	return r.Name
}
func (r UsersContactPostVO) GetSurname() string {
	return r.Surname
}
func (r UsersContactPostVO) GetPhone() string {
	return r.Phone
}
func (r UsersContactPostVO) GetEmail() string {
	return r.Email
}
