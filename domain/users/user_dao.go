package users

import (
	"fmt"
	"github.com/eremitic/bookstore_users-api/datasources/mysql/users_db"
	"github.com/eremitic/bookstore_users-api/utils/date_utils"
	"github.com/eremitic/bookstore_users-api/utils/errors"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsert      = "INSERT INTO users(first_name,last_name,email,date_created)VALUES(?,?,?,?)"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundErr(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsert)
	if err != nil {
		return errors.NewInternalErr("user insert q err")
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadReqErr("email taken")
		}
		return errors.NewInternalErr("user save err")
	}

	userId, err := insRes.LastInsertId()

	if err != nil {
		return errors.NewInternalErr("user save err")
	}

	user.Id = userId

	return nil
}
