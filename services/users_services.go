package services

import (
	"github.com/eremitic/bookstore_users-api/domain/users"
	"github.com/eremitic/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if createErr := user.Validate(); createErr != nil {
		return nil, createErr
	}

	if createErr := user.Save(); createErr != nil {
		return nil, createErr
	}
	return &user, nil
}

func GetUser(id int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: id}

	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}
