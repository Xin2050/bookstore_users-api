package app

import (
	pingController "bookstore_users-api/src/controllers/ping"
	usersController "bookstore_users-api/src/controllers/users"
)

func mapUrls() {
	// ping
	router.GET("/ping", pingController.Ping)
	// users
	router.GET("/users/:user_id", usersController.GetUser)
	router.GET("/users/search", usersController.SearchUser)
	router.POST("/users", usersController.CreateUser)
}
