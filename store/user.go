package store

import (
	"github.com/shubhamku044/management/model"
	"github.com/shubhamku044/management/util"
)

func (store Postgres) CreateUser(user *model.User) error {
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "Creating new user", nil)
	response := store.DB.Create(user)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.CreateUser, "Error in creating new user", response.Error)
		return response.Error
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "New user created", user)
	return nil
}
