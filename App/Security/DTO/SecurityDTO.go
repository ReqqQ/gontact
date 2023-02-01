package AppSecuritySecurityDTO

type TokenUserDTO struct {
	Id      int    `validate:"required"`
	Name    string `validate:"required"`
	Surname string `validate:"required"`
}

func (r TokenUserDTO) GetId() int {
	return r.Id
}
