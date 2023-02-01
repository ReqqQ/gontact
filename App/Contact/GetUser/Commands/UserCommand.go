package AppContactGetUserCommand

import "log"

type UserCommand struct {
	UserId *int `params:"userId"`
}

func (r UserCommand) Validate() UserCommand {
	if r.UserId == nil {
		log.Fatal("User id cannot be blank")
	}

	return r
}

func (r UserCommand) GetUserId() int {
	return *r.UserId
}
