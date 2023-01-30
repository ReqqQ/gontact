package groupsDatabase

import (
	GroupsModels "gontact/API/groups/models"
	"gontact/database"
)

const (
	searchUserGroupTypeByUserId = "select gt.id,gt.name from groups_type gt where created_by = @UserId"
)

func GetUserGroups(vo GroupsModels.GroupTypeVO) []GroupsModels.GroupType {
	var groupTypeCollection []GroupsModels.GroupType
	var groupStruct GroupsModels.GroupType
	rows, _ := database.DB.Raw(searchUserGroupTypeByUserId, vo).Rows()

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&groupStruct.Id, &groupStruct.Name)
		groupTypeCollection = append(groupTypeCollection, groupStruct)
	}

	return groupTypeCollection
}
