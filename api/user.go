package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/management/model"
	"github.com/shubhamku044/management/util"
)

func (api ApiRoutes) UserRoutes(routes *gin.Engine) {
	group := routes.Group("/user")

	{
		group.POST("/create", api.CreateUser)
		group.GET("/list", api.GetUsers)
		group.GET("/:id", api.GetUser)
	}
}

// Handler to create a user
// @router /user/create [post]
// @summary Create a user
// @tags users
// @accept json
// @produce json
// @param user body model.User true "User object"
// @success 201 {object} model.User
func (api ApiRoutes) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "Creating new user", nil)
	api.Server.CreateUser(ctx)
}

// Handler to get all users
// @router /user/list [get]
// @summary Get all users
// @tags users
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.User
// @Security ApiKeyAuth
func (api ApiRoutes) GetUsers(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUsers, "Fetching all users", nil)
	api.Server.GetUsers(ctx)
}

// Handler to get a user by ID
// @router /user/{id} [get]
// @summary Get a user by ID
// @tags users
// @produce json
// @param id path string true "User ID"
// @success 200 {object} model.User
// @Security ApiKeyAuth
func (api ApiRoutes) GetUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.GetUser, "Fetching user", nil)
	api.Server.GetUser(ctx)
}
