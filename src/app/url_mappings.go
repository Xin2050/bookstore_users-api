package app

import (
	pingController "bookstore_users-api/src/controllers/ping"
	usersController "bookstore_users-api/src/controllers/users"
)

func mapUrls() {
	// ping
	router.GET("/ping", pingController.Ping)
	// users

	router.GET("/users/search", usersController.SearchUser)
	router.POST("/users", usersController.CreateUser)
	router.GET("/users/:user_id", usersController.GetUser)
	router.PUT("/users/:user_id", usersController.UpdateUser)
	router.PATCH("/users/:user_id", usersController.UpdateUser)
	router.DELETE("/users/:user_id", usersController.DeleteUser)

}
