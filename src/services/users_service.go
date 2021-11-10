package services

import (
	"bookstore_users-api/src/domain/users"
	"bookstore_users-api/src/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
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
func DeleteUser(user users.User) *errors.RestError {
	_, err := GetUser(user.Id)
	if err != nil {
		return err
	}
	delErr := user.Delete()
	if delErr != nil {
		return delErr
	}
	return nil
}
