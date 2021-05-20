package users

import (
	"fmt"

	"github.com/kirankothule/bookstore-users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) GetUser() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User not found with ID: %d", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s id already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User with user ID %d, already exist", user.ID))
	}
	userDB[user.ID] = user
	return nil
}
