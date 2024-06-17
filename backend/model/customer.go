package model

import (
	"myapp/backend/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	UUID    uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Address string    `json:"address"`
}

type CustomerFilter struct {
	Id       uint      `query:"id"`
	UUID     uuid.UUID `query:"uuid"`
	Name     string    `query:"name"`
	Age      int       `query:"age"`
	Address  string    `query:"address"`
	SortName bool      `query:"sortName"`
}

type CustomerDto struct {
	UUID       uuid.UUID        `json:"uuid"`
	Name       string           `json:"name"`
	Age        int              `json:"age"`
	Address    string           `json:"address"`
	Pagination utils.Pagination `json:"pagination"`
}

func ConvertDtoToModel(input CustomerDto) Customer {
	return Customer{
		UUID:    input.UUID,
		Name:    input.Name,
		Age:     input.Age,
		Address: input.Address,
	}
}
