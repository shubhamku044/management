package store

import (
	"github.com/google/uuid"
	"github.com/shubhamku044/management/model"
	"github.com/shubhamku044/management/util"
	"gorm.io/gorm"
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

func (store Postgres) GetUsers() ([]model.User, error) {
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "Creating new user", nil)
	users := []model.User{}
	response := store.DB.Find(&users)
	if response.Error != nil {
		util.Log(model.LogLevelError, model.StorePackage, model.CreateUser, "Error in fetching all users", response.Error)
		return users, response.Error
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "List of all users", users)
	return users, nil
}

func (store Postgres) GetUser(userId uuid.UUID) (model.User, error) {

	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "Fetching records of user from DB", nil)
	user := model.User{}
	response := store.DB.First(&user, "id = ?", userId)
	if response.Error != nil {
		if response.Error == gorm.ErrRecordNotFound {
			util.Log(model.LogLevelError, model.StorePackage, model.CreateUser, "User not found", response.Error)
			return user, response.Error
		}
		util.Log(model.LogLevelError, model.StorePackage, model.CreateUser, "Error in fetching user", response.Error)

		return user, response.Error
	}
	util.Log(model.LogLevelInfo, model.StorePackage, model.CreateUser, "User fetched", user)
	return user, nil
}
