package model

import (
	"myapp/backend/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UUID          uuid.UUID       `json:"uuid"`
	Name          string          `json:"name"`
	Code          string          `json:"code"`
	ProductDetail []ProductDetail `gorm:"foreignKey:ProductId" json:"productDetail"`
}

type ProductFilter struct {
	Id       uint      `query:"id"`
	UUID     uuid.UUID `query:"uuid"`
	Name     string    `query:"name"`
	Code     string    `query:"code"`
	SortName bool      `query:"sortName"`
}

type ProductDto struct {
	UUID          uuid.UUID          `json:"uuid"`
	Name          string             `json:"name"`
	Code          string             `json:"code"`
	Pagination    utils.Pagination   `json:"pagination"`
	ProductDetail []ProductDetailDto `json:"productDetail"`
}

func ConvertProductDtoToModel(input ProductDto) Product {
	return Product{
		UUID:          input.UUID,
		Name:          input.Name,
		Code:          input.Code,
		ProductDetail: ConvertProductDetailDtosToModel(input.ProductDetail),
	}
}

func ConvertProductDetailDtosToModel(input []ProductDetailDto) []ProductDetail {
	var details []ProductDetail
	for v := range input {
		details = append(details, ConvertProductDetailDtoToModel(input[v]))
	}
	return details
}
