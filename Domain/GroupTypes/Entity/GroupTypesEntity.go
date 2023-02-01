package DomainGroupTypesEntity

type GroupTypesEntity struct {
	Id        int
	Name      string
	CreatedBy string
}

func (r GroupTypesEntity) GetId() int {
	return r.Id
}
func (r GroupTypesEntity) GetName() string {
	return r.Name
}
func (r GroupTypesEntity) GetCreatedBy() string {
	return r.CreatedBy
}
