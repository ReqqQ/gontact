package DomainUsers

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/matthewhartstonge/argon2"
	"github.com/thanhpk/randstr"
	DomainUsersEntity "gontact/Domain/Users/Entity"
	DomainUsersVO "gontact/Domain/Users/VO"
	InterfaceUsers "gontact/Interface/Users"
)

func GetUser(vo DomainUsersVO.UserVO) DomainUsersEntity.UsersEntity {
	return getUserEntity(InterfaceUsers.GetDbUser(vo, false))
}
func GetUserByToken(vo DomainUsersVO.UserTokenVO) DomainUsersEntity.UsersEntity {
	return getUserEntity(InterfaceUsers.GetDbUser(vo, true))
}
func GetUserContacts(vo DomainUsersVO.UserContactVO) []DomainUsersEntity.UsersContacts {
	return getUserContactsCollection(InterfaceUsers.GetDbUserContacts(vo))
}
func GetUsers() {
	return
}
func CreateUserContact(entity DomainUsersEntity.UsersContacts) {
	InterfaceUsers.CreateUserContact(entity)
}
func IsUserContactExists(entity DomainUsersEntity.UsersContacts, collection []DomainUsersEntity.UsersContacts) bool {
	isExists := false
	for _, dbEntity := range collection {
		if entity.GetEmail() == dbEntity.GetEmail() {
			isExists = true
			break
		}
	}

	return isExists
}
func HashPassword(password string) string {
	argon := argon2.DefaultConfig()
	hash, _ := argon.Hash([]byte(password), nil)

	token := string(hash.Encode())

	return token
}
func GenerateUserToken() string {
	argon := argon2.DefaultConfig()
	hash, _ := argon.Hash([]byte(randstr.Hex(5)), []byte(strconv.FormatInt(time.Now().UnixMicro(), 10)))

	token := hex.EncodeToString(hash.Hash)

	return token
}
