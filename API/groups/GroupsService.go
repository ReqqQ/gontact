package groups

import (
	groupsDatabase "gontact/API/groups/database"
	GroupsModels "gontact/API/groups/models"
)

func GetUserGroupTypes(vo GroupsModels.GroupTypeVO) []GroupsModels.GroupType {
	return groupsDatabase.GetUserGroups(vo)
}
