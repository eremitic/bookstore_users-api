package users

import (
	"fmt"
	"github.com/eremitic/bookstore_users-api/datasources/mysql/users_db"
	"github.com/eremitic/bookstore_users-api/logger"
	"github.com/eremitic/bookstore_users-api/utils/errors"
)

const (
	queryInsert       = "INSERT INTO users(first_name,last_name,email,date_created,password,status)VALUES(?,?,?,?,?,?)"
	queryUpdate       = "UPDATE users SET first_name=?,last_name=?,email=? where id=?"
	queryDelete       = "DELETE FROM users where id=?"
	queryGet          = "SELECT id,first_name,last_name,email,date_created,status from users where id=?"
	queryFindByStatus = "SELEC id,first_name,last_name,email,date_created,status from users where status=?"
)

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGet)

	if err != nil {
		logger.Error("user get prepare query err", err)
		return errors.NewInternalErr("db err")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("user get query err", err)
		return errors.NewInternalErr("db err")

	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsert)
	if err != nil {
		logger.Error("user prepare q query err", err)
		return errors.NewInternalErr("db err")
	}
	defer stmt.Close()

	insRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)

	if err != nil {
		logger.Error("user insert q query err", err)
		return errors.NewInternalErr("db err")

	}

	userId, err := insRes.LastInsertId()

	if err != nil {
		logger.Error("user save  query err", err)
		return errors.NewInternalErr("user save err")
	}

	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdate)
	if err != nil {
		logger.Error("user update prepare query err", err)
		return errors.NewInternalErr("user update q db err")
	}
	defer stmt.Close()

	upRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		logger.Error("user update query err", err)
		return errors.NewInternalErr("user update q db err")
	}

	_, err = upRes.RowsAffected()

	if err != nil {
		logger.Error("user update query row affected err", err)
		return errors.NewInternalErr("user update err")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryDelete)
	if err != nil {
		logger.Error("user delete prepare query err", err)
		return errors.NewInternalErr("user delete prepare q err")
	}
	defer stmt.Close()

	upRes, err := stmt.Exec(user.Id)

	if err != nil {
		logger.Error("user delete query err", err)
		return errors.NewInternalErr("user delete db")
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
		logger.Error("user find status prepare query err", err)
		return nil, errors.NewInternalErr("db err")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		logger.Error("user find status query err", err)
		return nil, errors.NewInternalErr("db err")
	}

	defer rows.Close()

	res := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("user find status parse query err", err)
			return nil, errors.NewInternalErr("db parse err")
		}
		res = append(res, user)
	}

	if len(res) == 0 {
		return nil, errors.NewNotFoundErr(fmt.Sprintf("no user match stauts %s", status))
	}
	return res, nil

}
