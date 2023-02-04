package AppContactCommands

type UserCommandInterface interface {
	GetUserId() int
	GetGroupId() *int
}

type GroupTypesCommand struct {
	UserId int `params:"userId"`
}

type UserCommand struct {
	UserId int
}

type UserContactsCommand struct {
	UserId  int
	GroupId *int `params:"groupId" validate:"omitempty,number,required"`
}

type CreateContactCommand struct {
	UserId  int
	GroupId *int
	Name    *string `params:"name" validate:"required,alpha"`
	Surname *string `params:"surname" validate:"required,alpha"`
	Email   *string `params:"email" validate:"required,email"`
	Phone   *string `params:"phone" validate:"required,e164"`
}

func (r CreateContactCommand) GetGroupId() *int {
	return r.GroupId
}

func (r CreateContactCommand) GetUserId() int {
	return r.UserId
}

func (r CreateContactCommand) GetUserName() string {
	return *r.Name
}

func (r CreateContactCommand) GetUserSurname() string {
	return *r.Surname
}

func (r CreateContactCommand) GetUserEmail() string {
	return *r.Email
}

func (r CreateContactCommand) GetUserPhone() string {
	return *r.Phone
}

func (r CreateContactCommand) SetUserId(userId int) CreateContactCommand {
	r.UserId = userId

	return r
}

func (r GroupTypesCommand) GetUserId() int {
	return r.UserId
}

func (r UserCommand) GetUserId() int {
	return r.UserId
}

func (r UserContactsCommand) SetUserId(userId int) UserContactsCommand {
	r.UserId = userId

	return r
}

func (r *UserContactsCommand) SetGroupId(groupId *int) {
	r.GroupId = groupId
}

func (r UserContactsCommand) GetUserId() int {
	return r.UserId
}

func (r UserContactsCommand) GetGroupId() *int {
	return r.GroupId
}

func (r UserContactsCommand) GetGroupIdAsValue() int {
	return *r.GroupId
}
