package app

import (
	"github.com/eremitic/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("start app")
	router.Run(":8080")
}
