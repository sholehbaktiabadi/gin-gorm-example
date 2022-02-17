package handler

import (
	"encoding/json"
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

func (r HandlerUserReciever) GetoneByEmail(ctx *gin.Context) {
	var user = user.User{}
	param := ctx.Param("email")
	user.Email = param
	res, err := r.user.GetoneByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) Getone(ctx *gin.Context) {
	token := ctx.MustGet("user")
	userByte, _ := json.Marshal(token)
	var jwt = helper.JwtMap{}
	err := json.Unmarshal(userByte, &jwt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}

	param := ctx.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := r.user.Getone(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) Update(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	var user = user.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	res, err := r.user.Update(user, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) Delete(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	res, err := r.user.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ResErr(http.StatusNotFound, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) GetAll(ctx *gin.Context) {
	res, err := r.user.Getall()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
}

func (r HandlerUserReciever) AdminUserRoutes(routerGroup *gin.RouterGroup) {
	m := middleware.Middleware{}
	routerGroup.GET("/email/:email", m.AuthenticationAdmin(), r.GetoneByEmail)
	routerGroup.GET("/:id", m.AuthenticationAdmin(), r.Getone)
	routerGroup.GET("/", m.AuthenticationAdmin(), r.GetAll)
	routerGroup.PUT("/:id", m.AuthenticationAdmin(), r.Update)
	routerGroup.DELETE("/:id", m.AuthenticationAdmin(), r.Delete)
}

func NewAdminUserHandler(user user.Init) HandlerUserReciever {
	return HandlerUserReciever{
		user: user,
	}
}
