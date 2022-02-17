package handler

import (
	"net/http"
	"v1/api/response"
	"v1/helper"
	"v1/middleware"
	"v1/repository/admin"
	"v1/repository/user"

	"github.com/gin-gonic/gin"
)

type AuthRecivier struct {
	user  user.Init
	admin admin.Init
}

func (r AuthRecivier) Register(ctx *gin.Context) {
	var (
		user user.User
	)

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	password, err := helper.GeneratePassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	user.Password = string(password)
	res, err := r.user.Create(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r AuthRecivier) Login(ctx *gin.Context) {
	var user = user.User{}
	m := middleware.Middleware{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	data, err := r.user.Login(user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, err.Error()))
		return
	}
	err = helper.CompareHashAndPassword([]byte(data.Password), user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, "password not match"))
		return
	}
	res, err := m.JwtSign(data.ID, false)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ResErr(http.StatusUnauthorized, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r AuthRecivier) LoginAdmin(ctx *gin.Context) {
	var admin = admin.Admin{}
	var m = middleware.Middleware{}
	err := ctx.ShouldBindJSON(&admin)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	data, err := r.admin.Login(admin)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, err.Error()))
		return
	}
	fail := helper.CompareHashAndPassword([]byte(data.Password), admin.Password)
	if fail != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, fail.Error()))
		return
	}
	res, err := m.JwtSign(data.ID, true)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r AuthRecivier) RegisterAdmin(ctx *gin.Context) {
	var admin = admin.Admin{}
	err := ctx.ShouldBindJSON(&admin)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	password, err := helper.GeneratePassword(admin.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	admin.Password = string(password)
	res, err := r.admin.Create(admin)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r AuthRecivier) UserAuthRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)
}

func (r AuthRecivier) AdminAuthRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/register", r.RegisterAdmin)
	routerGroup.POST("/login", r.LoginAdmin)
}

func NewAuthHandler(user user.Init, admin admin.Init) AuthRecivier {
	return AuthRecivier{
		user:  user,
		admin: admin,
	}
}
