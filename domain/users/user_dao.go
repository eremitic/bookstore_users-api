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
	queryGet         = "SELECT * from users where id=?"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGet)

	if err != nil {
		return errors.NewInternalErr("user get q err")
	}

	defer stmt.Close()
	reqId := user.Id
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return errors.NewNotFoundErr(fmt.Sprintf("user %d not found", reqId))
	}

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
