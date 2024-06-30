package user

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserDto struct {
	Id          uint
	UUID        uuid.UUID
	Email       string
	Password    string
	PhoneNumber string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name    string 
}