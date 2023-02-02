package AppContactGetUserCommand

type UserCommand struct {
	UserId int
}

func (r UserCommand) GetUserId() int {
	return r.UserId
}
