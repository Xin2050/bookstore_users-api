package users

import (
	"bookstore_users-api/src/datasources/mysql/users_db"
	"bookstore_users-api/src/logger"
	"bookstore_users-api/src/utils/errors"
	"fmt"
)

const (
	queryInsertUser       = "Insert into users(firstName, lastName, email, password, status, dateCreated) values (?,?,?,?,?,?);"
	queryGetUser          = "Select id, firstName, lastName, email, dateCreated,status from users where id=?;"
	queryUpdateUser       = "update users set firstName=?, lastName=?, email=? where id=?;"
	queryDeleteUser       = "delete from users where id =?;"
	queryGetUsersByStatus = "Select id, firstName, lastName, email, dateCreated, status from users where status=?;"
)

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	if err != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.InternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.InternalServerError("database error")
	}

	defer stmt.Close()
	insertRs, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email,
		user.Password, user.Status, user.DateCreated)
	if saveErr != nil {
		logger.Error("error when trying to save user statement", err)
		return errors.InternalServerError("database error")
	}
	userId, err := insertRs.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get user id after saved user", err)
		return errors.InternalServerError("database error")
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		logger.Error("error when trying to update user statement", err)
		return errors.InternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.InternalServerError("database error")
	}
	defer stmt.Close()
	_, delErr := stmt.Exec(user.Id)
	if delErr != nil {
		logger.Error("error when trying to delete user statement", err)
		return errors.InternalServerError("database error")
	}
	return nil
}

func (user *User) FindUsersByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryGetUsersByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by id statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by id statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.DateCreated,
			&user.Status,
		); err != nil {
			logger.Error("error when trying to get user field from search results", err)
			return nil, errors.InternalServerError("database error")
		}
		results = append(results, user)
	}
	// if no rows found
	if len(results) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
