package user

import uuid "github.com/satori/go.uuid"

type UserFilter struct {
	Id          uint
	UUID        uuid.UUID `query:"uuid"`
	Email       string    `query:"email"`
	PhoneNumber string    `query:"phoneNumber"`
	Name        string    `query:"name"`
	SortByName bool `query:"sortByName"`
}
