package AppContactGetUserContactsCommand

import (
	"log"
	"regexp"
	"strconv"
)

type UserContactsCommand struct {
	UserId  *int `params:"userId"`
	GroupId *int `params:"groupId"`
}

func (r UserContactsCommand) Validate() UserContactsCommand {
	if r.UserId == nil {
		log.Fatal("User id cannot be blank")
	}

	if r.GroupId != nil && !regexp.MustCompile(`\d`).MatchString(r.GetGroupIdAsValue()) {
		log.Fatal("Group Id must be int")
	}
	return r
}
func (r UserContactsCommand) GetUserId() int {
	return *r.UserId
}
func (r UserContactsCommand) GetGroupId() *int {
	return r.GroupId
}
func (r UserContactsCommand) GetGroupIdAsValue() string {
	return strconv.Itoa(*r.GroupId)
}
