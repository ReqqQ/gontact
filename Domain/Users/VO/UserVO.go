package DomainUsersVO

type UserVO struct {
	UserId int
}
type UserContactVO struct {
	UserId  int
	GroupId *int
}

func (r UserVO) GetUserId() int {
	return r.UserId
}
func (r UserContactVO) GetUserId() int {
	return r.UserId
}
func (r UserContactVO) GetGroupId() *int {
	return r.GroupId
}
