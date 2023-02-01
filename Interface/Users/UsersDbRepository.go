package InterfaceUsers

import (
	"database/sql"

	AppDatabase "gontact/App/Database"
	DomainUsersVO "gontact/Domain/Users/VO"
)

const (
	queryAnd                  = " and"
	queryWhere                = " where"
	searchById                = "select id,name,surname from users where id = @UserId"
	searchByToken             = "select id,name,surname from users where token = @Token"
	searchUserContactsByQuery = "select uc.id,uc.user_id,uc.name,uc.surname,uc.email," +
		"uc.phone from users_contacts uc "
	searchUserContactsByUserId  = " uc.user_id = @UserId"
	searchUserContactsByGroupId = " cg.group_id = @GroupId"
	searchUserContactsJoinGroup = " left join contacts_group cg on cg.user_contact_id=uc.id"
)

func GetDbUser(vo interface{}, isToken bool) *sql.Row {
	return AppDatabase.DB.Raw(getUserDbQuery(isToken), vo).Row()
}
func getUserDbQuery(isToken bool) string {
	query := searchById
	if isToken {
		query = searchByToken
	}

	return query
}
func GetDbUserContacts(vo DomainUsersVO.UserContactVO) *sql.Rows {
	rows, _ := AppDatabase.DB.Raw(getUserContactQuery(vo), vo).Rows()

	return rows
}
func getUserContactQuery(vo DomainUsersVO.UserContactVO) string {
	sqlQuery := searchUserContactsByQuery
	if vo.GetGroupId() != nil {
		sqlQuery += searchUserContactsJoinGroup + queryWhere + searchUserContactsByGroupId + queryAnd + searchUserContactsByUserId
	} else {
		sqlQuery += queryWhere + searchUserContactsByUserId
	}

	return sqlQuery
}
