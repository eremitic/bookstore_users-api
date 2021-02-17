package mysql_utils

import (
	"github.com/eremitic/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundErr("no record given id")
		}

		return errors.NewInternalErr("err parsing db response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadReqErr("invalid data")
	}
	return errors.NewInternalErr("err process db req")

}