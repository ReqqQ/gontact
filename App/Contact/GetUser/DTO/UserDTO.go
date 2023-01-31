package AppContactGetUserDTO

type UserDTO struct {
	Id      int    `json:"userId,omitempty"`
	Name    string `json:"userName,omitempty"`
	Surname string `json:"userSurname,omitempty"`
}
