package users

import (
	"fmt"

	"github.com/kirankothule/bookstore-users-api/datasources/mysql/users_db"
	"github.com/kirankothule/bookstore-users-api/utils/date_utils"
	"github.com/kirankothule/bookstore-users-api/utils/errors"
	"github.com/kirankothule/bookstore-users-api/utils/mysql_utils"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created from users WHERE id=?;"
	indexUniqueEmail = "email_UNIQUE"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName,
		&user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	user.DateCreated = date_utils.GetDateNowString()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userID, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf(" Error while trying to save user: %s", err.Error()))
	}
	user.ID = userID
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
