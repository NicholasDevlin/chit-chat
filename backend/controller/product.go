package controller

import (
	"myapp/backend/model"
	"myapp/backend/service"
	"myapp/backend/utils"

	"github.com/labstack/echo/v4"
)

type productController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *productController {
	return &productController{productService}
}

func (c *productController) CreateProduct(e echo.Context) error {
	var pagination utils.Pagination
	var input model.ProductDto
	e.Bind(&input)
	Product := model.ConvertProductDtoToModel(input)
	pagination = input.Pagination

	var err error
	err, pagination.CurrentPage = c.productService.SaveProduct(&Product)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}

	return utils.NewSuccessPaginationResponse(e, Product, pagination)
}

func (c *productController) GetAllProduct(e echo.Context) error {
	var filter model.ProductFilter
	if err := e.Bind(&filter); err != nil {
		return err
	}
	var pagination utils.Pagination
	if err := e.Bind(&pagination); err != nil {
		return err
	}

	res, err := c.productService.GetAllProduct(filter, &pagination)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}

	return utils.NewSuccessPaginationResponse(e, res, pagination)
}