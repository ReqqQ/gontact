package AppContactGetGroupTypesCommand

type GroupTypesCommand struct {
	UserId int `params:"userId"`
}

func (r GroupTypesCommand) GetUserId() int {
	return r.UserId
}
