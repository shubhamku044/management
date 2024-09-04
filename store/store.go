package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postges struct {
	DB *gorm.DB
}

func (store *Postges) NewStore() error {
	dsn := "host=localhost user=management password=management dbname=management port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	store.DB = db
	return nil
}
