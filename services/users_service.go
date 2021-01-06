package services

import (
	"github.com/said-saifi/users-mvc/domain/users"
	"github.com/said-saifi/users-mvc/utils/errors"
)

func GetUser(id int64) (*users.User, *errors.RestErr) {
	return users.Get(id)
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
