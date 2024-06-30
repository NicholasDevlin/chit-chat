package user

import uuid "github.com/satori/go.uuid"

type UserRes struct {
	UUID        uuid.UUID `json:"uuid" form:"uuid"`
	Email       string    `json:"email" form:"email"`
	Token       string    `json:"token" form:"token"`
	PhoneNumber string    `json:"phoneNumber" form:"phoneNumber"`
	Name        string    `json:"name" form:"name"`
}
