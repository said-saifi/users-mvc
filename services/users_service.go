package services

import (
	"github.com/said-saifi/users-mvc/domain/users"
	"github.com/said-saifi/users-mvc/utils/errors"
)

//TODO: save object in db
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
