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
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	curerntUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			curerntUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			curerntUser.LastName = user.LastName
		}
		if user.Email != "" {
			curerntUser.Email = user.Email
		}
	} else {
		curerntUser.FirstName = user.FirstName
		curerntUser.LastName = user.LastName
		curerntUser.Email = user.Email
	}
	if err := curerntUser.Update(); err != nil {
		return nil, err
	}
	return curerntUser, nil
}
