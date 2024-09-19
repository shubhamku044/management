package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      Name      `gorm:"embedded" json:"name" binding:"required"`
	Address   Address   `gorm:"embedded" json:"address" binding:"required"`
	Active    bool      `json:"active" example:"true"`
	Email     string    `json:"email" gorm:"unique; not null" example:"test@gmail.com" binding:"required"`
	Password  string    `json:"password" gorm:"not null" example:"password" binding:"required"`
	CreatedBy string    `json:"created_by" example:"shubham" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedBy string    `json:"deleted_by"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Name struct {
	FirstName  string `json:"first_name" example:"John" binding:"required"`
	MiddleName string `json:"middle_name" example:"A." binding:"required"`
	LastName   string `json:"last_name" example:"Doe" binding:"required"`
}

type Address struct {
	Lane     string `json:"lane" example:"123 Main St."`
	Village  string `json:"village" example:"Springfield"`
	City     string `json:"city" example:"Springfield" binding:"required"`
	District string `json:"district" example:"Greene" binding:"required"`
	Pincode  int    `json:"pincode" example:"12345" binding:"required"`
	State    string `json:"state" example:"IL" binding:"required"`
}
