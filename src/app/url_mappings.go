package app

import (
	pingController "bookstore_users-api/src/controllers/ping"
	usersController "bookstore_users-api/src/controllers/users"
	"bookstore_users-api/src/logger"
)

func mapUrls() {
	// ping
	router.GET("/ping", pingController.Ping)
	// users

	router.POST("/users", usersController.CreateUser)
	router.GET("/users/:user_id", usersController.GetUser)
	router.PUT("/users/:user_id", usersController.UpdateUser)
	router.PATCH("/users/:user_id", usersController.UpdateUser)
	router.DELETE("/users/:user_id", usersController.DeleteUser)

	router.GET("/internal/users/search", usersController.FindUsersByStatus)
	logger.Info("Routers were loaded.")
}
