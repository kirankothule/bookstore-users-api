package app

import (
	"github.com/gin-gonic/gin"

	"github.com/kirankothule/bookstore-users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start app....")
	router.Run(":8081")
}
