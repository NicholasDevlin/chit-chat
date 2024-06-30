package user

import uuid "github.com/satori/go.uuid"

type UserReq struct {
	UUID        uuid.UUID `json:"uuid" form:"uuid"`
	Email       string    `json:"email" form:"email"`
	Password    string    `json:"password" form:"password"`
	PhoneNumber string    `json:"phoneNumber" form:"phoneNumber"`
	Name        string    `json:"name" form:"name"`
}
