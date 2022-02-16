package api

import (
	"v1/api/handler"
	"v1/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(gorm gorm.DB) *gin.Engine {
	var (
		router      = gin.New()
		userInit    = user.NewUser(&gorm)
		userHandler = handler.NewUserHandler(userInit)
	)
	userHandler.Mount(router.Group(("user")))
	return router
}
