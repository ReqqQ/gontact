package AppContactGetGroupTypesDTO

type GroupTypesDTO struct {
	Id        int
	Name      string
	CreatedBy string `json:"-"`
}

func (r GroupTypesDTO) GetId() int {
	return r.Id
}
func (r GroupTypesDTO) GetName() string {
	return r.Name
}
func (r GroupTypesDTO) GetCreatedBy() string {
	return r.CreatedBy
}
