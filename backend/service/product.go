package service

import (
	"myapp/backend/model"
	"myapp/backend/repositories"
	"myapp/backend/utils"
	"myapp/backend/utils/errors"

	uuid "github.com/satori/go.uuid"
)

type IProductService interface {
	GetAllProduct(filter model.ProductFilter, pagination *utils.Pagination) ([]model.Product, error)
	GetProduct(filter model.ProductFilter) (model.Product, error)
	SaveProduct(input *model.Product) (error, int)
	// DeleteTransaction(id uuid.UUID) (transaction.TransactionRes, error)
}

type productService struct {
	productRepository repositories.IProductRpository
}

// GetAllProduct implements IProductService.
func (c *productService) GetAllProduct(filter model.ProductFilter, pagination *utils.Pagination) ([]model.Product, error) {
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}
	if pagination.CurrentPage == 0 {
		pagination.CurrentPage = 1
	}
	res, err := c.productRepository.GetAllProduct(filter, pagination)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetProduct implements IProductService.
func (*productService) GetProduct(filter model.ProductFilter) (model.Product, error) {
	panic("unimplemented")
}

// SaveProduct implements IProductService.
func (c *productService) SaveProduct(input *model.Product) (error, int) {
	var err error
	if input.Name == "" {
		return errors.ERR_NAME_IS_EMPTY, 1
	}
	var newData model.Product
	newData = *input
	if input.UUID != uuid.Nil {
		*input, err = c.productRepository.GetProduct(&model.ProductFilter{UUID: newData.UUID})
		if err != nil {
			return errors.ERR_CREATE_PRODUCT,1 
		}
		input.Code = newData.Code
		input.Name = newData.Name
		for _,v := range newData.ProductDetail {
			detail, err := c.productRepository.GetProductDetail(&model.ProductDetailFilter{UUID: v.UUID})
			if err != nil {
				return errors.ERR_CREATE_PRODUCT,1
			}
			detail.Price = v.Price
			detail.Size = v.Size
			input.ProductDetail = append(input.ProductDetail, detail)
		}
	}

	err, pageNumber := c.productRepository.SaveProduct(input)
	if err != nil {
		return errors.ERR_CREATE_PRODUCT, 1
	}
	return err, pageNumber
}

func NewProductService(repo repositories.IProductRpository) IProductService {
	return &productService{
		productRepository: repo,
	}
}
