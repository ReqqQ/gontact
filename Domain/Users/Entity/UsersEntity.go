package DomainUsersEntity

type UsersEntity struct {
	Id      int
	Name    string
	Surname string
}

type UserContacts struct {
	Id      int
	UserId  int
	Name    string
	Surname string
	Email   string
	Phone   string
}

func (r UserContacts) GetId() int {
	return r.Id
}
func (r UserContacts) GetUserId() int {
	return r.UserId
}
func (r UserContacts) GetName() string {
	return r.Name
}
func (r UserContacts) GetSurname() string {
	return r.Surname
}
func (r UserContacts) GetEmail() string {
	return r.Email
}
func (r UserContacts) GetPhone() string {
	return r.Phone
}

func (r UsersEntity) GetId() int {
	return r.Id
}
func (r UsersEntity) GetName() string {
	return r.Name
}
func (r UsersEntity) GetSurname() string {
	return r.Surname
}
