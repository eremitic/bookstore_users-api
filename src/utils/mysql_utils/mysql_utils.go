package mysql_utils

import (
	"fmt"
	"github.com/eremitic/bookstore_users-api/src/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundErr("no record given id")
		}
		fmt.Println(err.Error())
		return errors.NewInternalErr("err parsing db response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadReqErr("invalid data")
	}
	return errors.NewInternalErr(sqlErr.Error())

}
