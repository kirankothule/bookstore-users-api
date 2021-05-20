package services

import (
	"github.com/kirankothule/bookstore-users-api/domain/users"
	"github.com/kirankothule/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{
		ID: userID,
	}
	if err := user.GetUser(); err != nil {
		return nil, err
	}
	return user, nil
}
