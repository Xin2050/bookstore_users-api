package users

import (
	"bookstore_users-api/src/utils/errors"
	"github.com/go-playground/validator/v10"
	"strings"
	"unicode"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" validate:"email"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
	Password    string `json:"password" validate:"min=8,max=50,passwd"`
}

type Users []User

var validatePassword = func(fl validator.FieldLevel) bool {
	var (
		noUpper   = true
		noLower   = true
		noNumber  = true
		noSpecial = true
	)
	for _, c := range fl.Field().String() {
		switch {
		case unicode.IsUpper(c):
			noUpper = false
		case unicode.IsLower(c):
			noLower = false
		case unicode.IsNumber(c):
			noNumber = false
		case unicode.IsPunct(c):
			noSpecial = false
		}
	}
	return !(noLower || noUpper || noNumber || noSpecial)
}

func (user *User) Validate() *errors.RestError {
	v := validator.New()

	_ = v.RegisterValidation("passwd", validatePassword)
	if err := v.Struct(user); err != nil {
		return errors.BadRequestError(err.Error())
	}
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	return nil
}
func (user *User) ValidateWithoutPassword() *errors.RestError {

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	return nil
}
