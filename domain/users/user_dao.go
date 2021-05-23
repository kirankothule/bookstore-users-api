package users

import (
	"fmt"
	"strings"

	"github.com/kirankothule/bookstore-users-api/datasources/mysql/users_db"
	"github.com/kirankothule/bookstore-users-api/utils/date_utils"
	"github.com/kirankothule/bookstore-users-api/utils/errors"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	indexUniqueEmail = "email_UNIQUE"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) GetUser() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	user.DateCreated = date_utils.GetDateNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("Mail id %s already registered", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf(
			"Error while saving user to DB: %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf(" Error while trying to save user: %s", err.Error()))
	}
	user.ID = userID
	return nil
}
