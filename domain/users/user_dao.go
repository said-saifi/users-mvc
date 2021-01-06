package users

import (
	"fmt"
	"time"

	"github.com/said-saifi/users-mvc/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func Get(id int64) (*User, *errors.RestErr) {
	res, ok := usersDB[id]
	if !ok {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user %d not found", id))
	}
	return res, nil

}

func (user *User) Save() *errors.RestErr {
	if _, ok := usersDB[user.Id]; ok {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.TsCreated = time.Now().Unix()
	usersDB[user.Id] = user
	return nil
}
