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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
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

func DeleteUser(id int64) *errors.RestErr {
	result := &users.User{Id: id}

	if err := result.Delete(); err != nil {
		return err
	}
	return nil

}
