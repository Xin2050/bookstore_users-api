package users

import (
	"bookstore_users-api/src/datasources/mysql/users_db"
	"bookstore_users-api/src/utils/errors"
	mysql_utils "bookstore_users-api/src/utils/mysql"
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
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	defer stmt.Close()
	insertRs, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email,
		user.Password, user.Status, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertRs.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()
	_, delErr := stmt.Exec(user.Id)
	if delErr != nil {
		return mysql_utils.ParseError(delErr)
	}
	return nil
}

func (user *User) FindUsersByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryGetUsersByStatus)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
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
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	// if no rows found
	if len(results) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
