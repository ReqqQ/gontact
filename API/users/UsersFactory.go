package users

import UserModels "gontact/API/users/models"

func CreateUserVO(id int, groupId *int) UserModels.UsersVO {
	return UserModels.UsersVO{
		UserId:  id,
		GroupId: groupId,
	}
}
