package store

import (
	"github.com/shubhamku044/management/model"
	"github.com/shubhamku044/management/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (store *Postgres) NewStore() error {
	dsn := "host=localhost user=management password=management dbname=management port=5432 sslmode=disable"
	util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, "Connecting to database", nil)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, "Error while connecting to database", err)
		return err
	}
	store.DB = db

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		util.Log(model.LogLevelInfo, model.StorePackage, model.NewStore, "Error while migrating the table", err)
	}

	return nil
}

type StoreOperations interface {
	NewStore() error
	CreateUser(user *model.User) error
}
