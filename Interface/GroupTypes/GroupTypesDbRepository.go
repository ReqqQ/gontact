package InterfaceGroupTypes

import (
	"database/sql"

	AppDatabase "gontact/App/Database"
	DomainGroupTypesVO "gontact/Domain/GroupTypes/VO"
)

const (
	searchById = "select id,name,created_by from groups_type where created_by = @UserId"
)

func GetDbGroupTypes(vo DomainGroupTypesVO.GroupTypesVO) *sql.Rows {
	rows, _ := AppDatabase.DB.Raw(searchById, vo).Rows()

	return rows
}
