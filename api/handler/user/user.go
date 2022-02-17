package handler

import (
	"encoding/json"
	"net/http"
	"v1/api/response"
	"v1/helper"
	"v1/middleware"
	"v1/repository/user"

	"github.com/gin-gonic/gin"
)

type HandlerUserReciever struct {
	user user.Init
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
	res, err := r.user.Getone(jwt.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) Update(ctx *gin.Context) {
	token := ctx.MustGet("user")
	userByte, _ := json.Marshal(token)
	var jwt = helper.JwtMap{}
	err := json.Unmarshal(userByte, &jwt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	var user = user.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResErr(http.StatusUnprocessableEntity, err.Error()))
		return
	}
	res, err := r.user.Update(user, jwt.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResErr(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResOK("success", res))
	return
}

func (r HandlerUserReciever) UserHandler(routerGroup *gin.RouterGroup) {
	m := middleware.Middleware{}
	routerGroup.GET("/", m.Authentication(), r.Getone)
	routerGroup.PUT("/", m.Authentication(), r.Update)
}

func NewUserHandler(user user.Init) HandlerUserReciever {
	return HandlerUserReciever{
		user: user,
	}
}
