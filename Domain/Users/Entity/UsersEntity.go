package DomainUsersEntity

type UsersEntity struct {
	Id      int
	Name    string
	Surname string
}

type UsersContacts struct {
	Id      int
	UserId  int
	Name    string
	Surname string
	Email   string
	Phone   string
}

func (r UsersContacts) GetId() int {
	return r.Id
}
func (r UsersContacts) GetUserId() int {
	return r.UserId
}
func (r UsersContacts) GetName() string {
	return r.Name
}
func (r UsersContacts) GetSurname() string {
	return r.Surname
}
func (r UsersContacts) GetEmail() string {
	return r.Email
}
func (r UsersContacts) GetPhone() string {
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
