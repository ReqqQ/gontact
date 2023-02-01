package DomainGroupTypesVO

type GroupTypesVO struct {
	UserId int
}

func (r GroupTypesVO) GetUserId() int {
	return r.UserId
}
