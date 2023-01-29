package users

import (
	UserModels "gontact/API/users/models"
	"gontact/database"
)

const (
	searchById                = "id = ?"
	searchUserContactsByQuery = "select u.id,u.user_id,u.name,u.surname,u.email,u.phone," +
		"gt.name from contacts_group cg left join groups_type gt on gt.id=cg." +
		"group_id left join users_contacts u on u.id=cg.user_contact_id where u.user_id = @UserId "
	searchUserByGroupId = "and cg.group_id = @GroupId"
)

func getDbUser(vo UserModels.UsersVO) UserModels.Users {
	var user UserModels.Users
	database.DB.Where(searchById, vo.GetId()).First(&user)

	return user
}
func getDbUserContacts(vo UserModels.UsersVO) []UserModels.UsersContacts {
	var usersContacts []UserModels.UsersContacts
	var userContactStruct UserModels.UsersContacts
	sqlQuery := searchUserContactsByQuery
	if vo.GetGroupId() != nil {
		sqlQuery += searchUserByGroupId
	}
	rows, _ := database.DB.Raw(sqlQuery, vo).Rows()

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&userContactStruct.Id, &userContactStruct.UserId, &userContactStruct.Name, &userContactStruct.Surname, &userContactStruct.Email, &userContactStruct.Phone, &userContactStruct.Group.Name)
		usersContacts = append(usersContacts, userContactStruct)
	}

	return usersContacts
}
