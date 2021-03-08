package services

import (
	"github.com/eremitic/bookstore_users-api/src/domain/users"
	"github.com/eremitic/bookstore_users-api/src/utils/crypto_utils"
	"github.com/eremitic/bookstore_users-api/src/utils/date_utils"
	"github.com/eremitic/bookstore_users-api/src/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if createErr := user.Validate(); createErr != nil {
		return nil, createErr
	}

	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMD5(user.Password)
	user.DateCreated = date_utils.GetDbString()
	if createErr := user.Save(); createErr != nil {
		return nil, createErr
	}
	return &user, nil
}

func (s *userService) GetUser(id int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: id}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email

	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userService) DeleteUser(id int64) *errors.RestErr {
	result := &users.User{Id: id}

	if err := result.Delete(); err != nil {
		return err
	}
	return nil

}

func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)

}