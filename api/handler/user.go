package handler

import (
	"net/http"
	"strconv"
	"v1/api/response"
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

func (r HandlerUserReciever) Mount(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", r.Register)
	routerGroup.GET("/email/:email", r.GetoneByEmail)
	routerGroup.GET("/:id", r.Getone)
	routerGroup.GET("/", r.GetAll)
	routerGroup.PUT("/:id", r.Update)
	routerGroup.DELETE("/:id", r.Delete)
}

func NewUserHandler(user user.Init) HandlerUserReciever {
	return HandlerUserReciever{
		user: user,
	}
}
