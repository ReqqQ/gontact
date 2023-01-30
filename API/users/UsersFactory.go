package users

import (
	"github.com/gofiber/fiber/v2"
	UserModels "gontact/API/users/models"
)

func CreateUserVO(id int, groupId *int) UserModels.UsersVO {
	return UserModels.UsersVO{
		UserId:  id,
		GroupId: groupId,
	}
}
func CreateUserContactPostVO(id int, c *fiber.Ctx) UserModels.UsersContactPostVO {
	vo := new(UserModels.UsersContactPostVO)
	c.BodyParser(vo)

	return vo.SetUserId(id)
}
func CreateUserContactStruct(vo UserModels.UsersContactPostVO) UserModels.UsersContacts {
	return UserModels.UsersContacts{
		UserId:  vo.GetUserId(),
		Name:    vo.GetName(),
		Surname: vo.GetSurname(),
		Phone:   vo.GetPhone(),
		Email:   vo.GetEmail(),
	}
}
