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
