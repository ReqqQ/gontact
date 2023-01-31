package AppContactGetGroupTypesCommand

import "log"

type GroupTypesCommand struct {
	UserId *int `params:"userId"`
}

func (r GroupTypesCommand) Validate() GroupTypesCommand {
	if r.UserId == nil {
		log.Fatal("User id cannot be blank")
	}

	return r
}

func (r GroupTypesCommand) GetUserId() int {
	return *r.UserId
}
