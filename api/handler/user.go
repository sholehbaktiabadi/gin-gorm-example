package handler

import (
	"net/http"
	"strconv"
	"v1/api/response"
	"v1/helper"
	"v1/middleware"
	"v1/user"

	"github.com/gin-gonic/gin"
)

type HandlerUserReciever struct {
	user user.Init
}

func (r HandlerUserReciever) Register(gin *gin.Context) {
	var (
		user user.User
	)

	err := gin.ShouldBindJSON(&user)
	if err != nil {
		gin.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	password, err := helper.GeneratePassword(user.Password)
	if err != nil {
		gin.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	user.Password = string(password)
	res, err := r.user.Create(user)
	if err != nil {
		gin.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	gin.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) GetoneByEmail(gin *gin.Context) {
	var user = user.User{}
	param := gin.Param("email")
	user.Email = param
	res, err := r.user.GetoneByEmail(user.Email)
	if err != nil {
		gin.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, err.Error()))
		return
	}
	gin.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) Getone(gin *gin.Context) {
	param := gin.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		gin.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := r.user.Getone(id)
	if err != nil {
		gin.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	gin.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) Update(gin *gin.Context) {
	param := gin.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		gin.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	var user = user.User{}
	if err := gin.ShouldBindJSON(&user); err != nil {
		gin.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	res, err := r.user.Update(user, id)
	if err != nil {
		gin.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	gin.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) Delete(gin *gin.Context) {
	param := gin.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		gin.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := r.user.Delete(id)
	if err != nil {
		gin.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, err.Error()))
		return
	}
	gin.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) GetAll(gin *gin.Context) {
	res, err := r.user.Getall()
	if err != nil {
		gin.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	gin.JSON(http.StatusOK, response.ResOK("success", res))
}

func (r HandlerUserReciever) Login(gin *gin.Context) {
	var user = user.User{}
	m := middleware.Middleware{}
	err := gin.ShouldBindJSON(&user)
	if err != nil {
		gin.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	data, err := r.user.Login(user)
	if err != nil {
		gin.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, err.Error()))
		return
	}
	err = helper.CompareHashAndPassword([]byte(data.Password), user.Password)
	if err != nil {
		gin.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, "password not match"))
		return
	}
	res, err := m.JwtSign(data.ID)
	if err != nil {
		gin.JSON(http.StatusUnauthorized, response.ResErr(http.StatusUnauthorized, err.Error()))
		return
	}
	gin.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) MountUser(routerGroup *gin.RouterGroup) {
	m := middleware.Middleware{}
	routerGroup.GET("/email/:email", m.Authentication(), r.GetoneByEmail)
	routerGroup.GET("/:id", m.Authentication(), r.Getone)
	routerGroup.GET("/", m.Authentication(), r.GetAll)
	routerGroup.PUT("/:id", m.Authentication(), r.Update)
	routerGroup.DELETE("/:id", m.Authentication(), r.Delete)
}

func (r HandlerUserReciever) MountAuth(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)
}

func NewUserHandler(user user.Init) HandlerUserReciever {
	return HandlerUserReciever{
		user: user,
	}
}
