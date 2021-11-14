package services

import (
	"bookstore_users-api/src/domain/users"
	crypto_utils "bookstore_users-api/src/utils/crypto"
	date_utils "bookstore_users-api/src/utils/date"
	"bookstore_users-api/src/utils/errors"
	mysql_utils "bookstore_users-api/src/utils/mysql"
)

const (
	StatusActive = "Active"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	FindUsersByStatus(string) (users.Users, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(users.User) *errors.RestError
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowForMySQL()
	user.Status = StatusActive
	hashedPwd, err := crypto_utils.HashPassword(user.Password)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	user.Password = hashedPwd
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.ValidateWithoutPassword(); err != nil {
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
		current.LastName = user.LastName
		current.FirstName = user.FirstName
		current.Email = user.Email
	}

	updateErr := current.Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return current, nil
}
func (s *usersService) DeleteUser(user users.User) *errors.RestError {
	_, err := s.GetUser(user.Id)
	if err != nil {
		return err
	}
	delErr := user.Delete()
	if delErr != nil {
		return delErr
	}
	return nil
}
func (s *usersService) FindUsersByStatus(status string) (users.Users, *errors.RestError) {
	user := &users.User{}
	return user.FindUsersByStatus(status)

}
