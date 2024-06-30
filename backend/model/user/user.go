package user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string
	Name        string
}
