package InterfaceUsers

import (
	"database/sql"

	AppDatabase "gontact/App/Database"
	DomainUsersEntity "gontact/Domain/Users/Entity"
	DomainUsersVO "gontact/Domain/Users/VO"
	"gorm.io/gorm"
)

const (
	searchById                = "select id,name,surname from users where id = @UserId"
	searchByToken             = "select id,name,surname from users where token = @Token"
	searchUserContactsByQuery = "select uc.id,uc.user_id,uc.name,uc.surname,uc.email," +
		"uc.phone from users_contacts uc "
	searchUserContactsByUserId  = "where uc.user_id = @UserId"
	searchUserContactsByGroupId = "left join contacts_group cg on cg.user_contact_id=uc.id where cg.group_id = @GroupId and uc.user_id = @UserId"
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
		sqlQuery += searchUserContactsByGroupId
	} else {
		sqlQuery += searchUserContactsByUserId
	}

	return sqlQuery
}
func CreateUserContact(entity DomainUsersEntity.UsersContacts) *gorm.DB {
	rest := AppDatabase.DB.Create(&entity)

	return rest
}
