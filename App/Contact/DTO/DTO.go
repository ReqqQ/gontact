package AppContactDTO

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

type UserContactDTO struct {
	Id      int
	UserId  int
	Name    string
	Surname string
	Email   string
	Phone   string
}
type UserDTO struct {
	Id      int    `json:"userId,omitempty"`
	Name    string `json:"userName,omitempty"`
	Surname string `json:"userSurname,omitempty"`
}
