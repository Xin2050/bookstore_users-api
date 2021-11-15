package app

import (
	"bookstore_users-api/src/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("application is Running on port 3000.")
	router.Run(":3000")
}
