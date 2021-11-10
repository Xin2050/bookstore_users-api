package mysql_utils

import (
	"bookstore_users-api/src/utils/errors"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) *errors.RestError {
	if err == sql.ErrNoRows {
		return errors.NotFoundError(fmt.Sprintf("no record found by given condictions"))
	}
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return errors.InternalServerError(fmt.Sprintf("error: %s", err))
	}
	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError(fmt.Sprintf("duplicated key: %s", sqlErr.Message))
	}
	return errors.InternalServerError(fmt.Sprintf("error processing request:%s", err))
}
