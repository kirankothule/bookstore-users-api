package services

import (
	"fmt"

	"github.com/kirankothule/bookstore-users-api/domain/users"
	"github.com/kirankothule/bookstore-users-api/utils/crypto_utils"
	"github.com/kirankothule/bookstore-users-api/utils/date_utils"
	"github.com/kirankothule/bookstore-users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	fmt.Println("in service")
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{
		ID: userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	curerntUser, err := s.GetUser(user.ID)
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
		if user.Status != "" {
			curerntUser.Status = user.Status
		}
	} else {
		curerntUser.FirstName = user.FirstName
		curerntUser.LastName = user.LastName
		curerntUser.Email = user.Email
		curerntUser.Status = user.Status
	}
	fmt.Println("upadte user: ", curerntUser)
	if err := curerntUser.Update(); err != nil {
		return nil, err
	}
	return curerntUser, nil
}

func (s *userService) DeleteUser(userID int64) *errors.RestErr {

	_, err := s.GetUser(userID)
	if err != nil {
		return err
	}
	user := &users.User{
		ID: userID,
	}
	return user.Delete()
}

func (s *userService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *userService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {

	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMD5(request.Password),
		Status:   users.StatusActive,
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
