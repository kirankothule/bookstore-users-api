package users

import (
	"fmt"
	"strings"

	"github.com/kirankothule/bookstore-users-api/datasources/mysql/users_db"
	"github.com/kirankothule/bookstore-users-api/logger"
	"github.com/kirankothule/bookstore-users-api/utils/errors"
	"github.com/kirankothule/bookstore-users-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status from users WHERE id=?;"
	indexUniqueEmail            = "email_UNIQUE"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, Last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, Last_name, email, date_created, status FROM users WHERE email=? and password=? and status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error while prearing Get user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName,
		&user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return errors.NewNotFoundError("User not found")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error while prearing insert user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {

		logger.Error("Error while saving the user into the database", err)
		return errors.NewInternalServerError("Database error")
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error while getting user id after user save", err)
		return errors.NewInternalServerError("Database error")
	}
	user.ID = userID
	return nil
}

func (user *User) Update() *errors.RestErr {
	fmt.Println("in update dao")
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error while prearing update user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	fmt.Println("update exec")
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.ID)
	if err != nil {
		logger.Error("Error while updating user", err)
		return errors.NewInternalServerError("Database update error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error while prearing dalete user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	if err != nil {
		logger.Error("Error while deleting user", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("Error while prearing find by status user statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error while getting using by status", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer rows.Close()

	result := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error while parsing resultset in findByStatus", err)
			return nil, errors.NewInternalServerError("Database error")
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.NewNotFoundError("No user found with given status")
	}
	return result, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("Error while prearing Get user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email, user.Password, user.Status)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName,
		&user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrNoRows) {
			return errors.NewNotFoundError("invalid user credentials")
		}
		return errors.NewNotFoundError("User not found")
	}
	return nil
}
