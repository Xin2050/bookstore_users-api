package users

import (
	"bookstore_users-api/src/datasources/mysql/users_db"
	date_utils "bookstore_users-api/src/utils/date"
	"bookstore_users-api/src/utils/errors"
	mysql_utils "bookstore_users-api/src/utils/mysql"
)

const (
	queryInsertUser = "Insert into users(firstName, lastName, email, dateCreated) values (?,?,?,?);"
	queryGetUser    = "Select id, firstName, lastName, email, dateCreated from users where id=?;"
	queryUpdateUser = "update users set firstName=?, lastName=?, email=? where id=?;"
	queryDeleteUser = "delete from users where id =?;"
)

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
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
	user.DateCreated = date_utils.GetNowForMySQL()
	defer stmt.Close()
	insertRs, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
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
