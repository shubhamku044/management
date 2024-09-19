package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shubhamku044/management/model"
	"github.com/shubhamku044/management/util"
)

func (server Server) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.CreateUser, "Creating new user", nil)
	user := model.User{}

	err := ctx.Bind(&user)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "Error in binding request", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	err = server.PostgresDB.CreateUser(&user)

	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.CreateUser, "Error in creating new user", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (server Server) GetUsers(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUsers, "Fetching all users", nil)
	users, err := server.PostgresDB.GetUsers()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUsers, "Error in fetching all users", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (server Server) GetUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ControllerPackage, model.GetUser, "Fetching user", nil)
	userId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUser, "Error in parsing user id", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user, err := server.PostgresDB.GetUser(userId)
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.GetUser, "Error in fetching user", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
