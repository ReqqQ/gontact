package AppContactGetUserContactsCommand

type UserContactsCommand struct {
	UserId  int
	GroupId *int `params:"groupId" validate:"omitempty,number,required"`
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
