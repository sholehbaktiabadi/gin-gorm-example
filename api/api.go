package api

import (
	adminhandler "v1/api/handler/admin"
	authhandler "v1/api/handler/auth"
	userhandler "v1/api/handler/user"
	"v1/repository/admin"
	"v1/repository/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(gorm gorm.DB) *gin.Engine {
	var (
		memberPrefix = "v1/"
		adminPrefix  = "v1/admin"
		router       = gin.New()
		userInit     = user.NewUser(&gorm)
		adminInit    = admin.NewAdmin(&gorm)
		adminHandler = adminhandler.NewAdminUserHandler(userInit)
		authHandler  = authhandler.NewAuthHandler(userInit, adminInit)
		userHandler  = userhandler.NewUserHandler(userInit)
	)

	// admin routes
	authHandler.AdminAuthRoutes(router.Group(adminPrefix + "/auth"))
	adminHandler.AdminUserRoutes(router.Group(adminPrefix + "/user"))

	// member routes
	authHandler.UserAuthRoutes(router.Group(memberPrefix + "/auth"))
	userHandler.UserHandler(router.Group(memberPrefix + "/user"))
	return router
}
