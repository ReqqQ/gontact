package groups

import GroupsModels "gontact/API/groups/models"

func GetGroupTypeVO(userId int) GroupsModels.GroupTypeVO {
	return GroupsModels.GroupTypeVO{UserId: userId}
}
