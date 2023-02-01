package DomainUsers

import (
	"database/sql"

	DomainUsersEntity "gontact/Domain/Users/Entity"
	DomainUsersVO "gontact/Domain/Users/VO"
)

func getUserEntity(row *sql.Row) DomainUsersEntity.UsersEntity {
	var userEntity DomainUsersEntity.UsersEntity

	row.Scan(&userEntity.Id, &userEntity.Name, &userEntity.Surname)

	return userEntity
}
func getUserContactsCollection(rows *sql.Rows) []DomainUsersEntity.UserContacts {
	var collection []DomainUsersEntity.UserContacts
	defer rows.Close()

	for rows.Next() {
		var userContactEntity DomainUsersEntity.UserContacts
		rows.Scan(&userContactEntity.Id, &userContactEntity.UserId, &userContactEntity.Name, &userContactEntity.Surname, &userContactEntity.Email, &userContactEntity.Phone)
		collection = append(collection, userContactEntity)
	}

	return collection
}
func GetUserTokenVO(userToken string) DomainUsersVO.UserTokenVO {
	return DomainUsersVO.UserTokenVO{
		Token: userToken,
	}
}
