package groupsModels

type GroupType struct {
	Id        int    `json:",omitempty"`
	Name      string `json:",omitempty"`
	CreatedBy int    `json:",omitempty"`
}
type GroupTypeVO struct {
	UserId int
}

func (gt GroupTypeVO) GetUserId() int {
	return gt.UserId
}
