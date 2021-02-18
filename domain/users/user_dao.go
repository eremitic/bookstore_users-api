package users

import (
	"fmt"
	"github.com/eremitic/bookstore_users-api/datasources/mysql/users_db"
	"github.com/eremitic/bookstore_users-api/utils/errors"
	"github.com/eremitic/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsert       = "INSERT INTO users(first_name,last_name,email,date_created,password,status)VALUES(?,?,?,?,?,?)"
	queryUpdate       = "UPDATE users SET first_name=?,last_name=?,email=? where id=?"
	queryDelete       = "DELETE FROM users where id=?"
	queryGet          = "SELECT id,first_name,last_name,email,date_created,status from users where id=?"
	queryFindByStatus = "SELECT id,first_name,last_name,email,date_created,status from users where status=?"
)

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGet)

	if err != nil {
		return errors.NewInternalErr("user get q err")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsert)
	if err != nil {
		return errors.NewInternalErr("user insert q err")
	}
	defer stmt.Close()

	insRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userId, err := insRes.LastInsertId()

	if err != nil {
		return errors.NewInternalErr("user save err")
	}

	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdate)
	if err != nil {
		return errors.NewInternalErr("user update q err")
	}
	defer stmt.Close()

	upRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	_, err = upRes.RowsAffected()

	if err != nil {
		return errors.NewInternalErr("user update err")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryDelete)
	if err != nil {
		return errors.NewInternalErr("user delete q err")
	}
	defer stmt.Close()

	upRes, err := stmt.Exec(user.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	_, err = upRes.RowsAffected()

	if err != nil {
		return errors.NewInternalErr("user delete err")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)

	if err != nil {
		return nil, errors.NewInternalErr(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		return nil, errors.NewInternalErr(err.Error())
	}

	defer rows.Close()

	res := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		res = append(res, user)
	}

	if len(res) == 0 {
		return nil, errors.NewNotFoundErr(fmt.Sprintf("no user match stauts %s", status))
	}
	return res, nil

}
