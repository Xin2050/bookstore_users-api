package users

import (
	"bookstore_users-api/src/utils/errors"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" validate:"email"`
	DateCreated string `json:"dateCreated"`
}

func (user *User) Validate() *errors.RestError {
	v := validator.New()
	if err := v.Struct(user); err != nil {
		return errors.BadRequestError(err.Error())
	}
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	return nil
}
