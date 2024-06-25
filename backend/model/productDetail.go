package model

import (
	"myapp/backend/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ProductDetail struct {
	gorm.Model
	UUID      uuid.UUID `json:"uuid"`
	Size      string    `json:"size"`
	Price     uint      `json:"price"`
	ProductId uint      `json:"productId"`
}
func (pd *ProductDetail) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID using the satori/go.uuid package
	pd.UUID = uuid.NewV4()
	return 
}

type ProductDetailFilter struct {
	Id        uint      `query:"id"`
	UUID      uuid.UUID `query:"uuid"`
	Size      string    `query:"size"`
	Price     uint      `query:"price"`
	ProductId uint      `query:"productId"`
}

type ProductDetailDto struct {
	UUID       uuid.UUID        `json:"uuid"`
	Size       string           `json:"size"`
	Price      uint             `json:"price"`
	ProductId  uint             `json:"productId"`
	Pagination utils.Pagination `json:"pagination"`
}

func ConvertProductDetailDtoToModel(input ProductDetailDto) ProductDetail {
	return ProductDetail{
		UUID:      input.UUID,
		Size:      input.Size,
		Price:     input.Price,
		ProductId: input.ProductId,
	}
}
