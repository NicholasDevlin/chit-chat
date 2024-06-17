package controller

import (
	"myapp/backend/model"
	"myapp/backend/service"
	"myapp/backend/utils"

	"github.com/labstack/echo/v4"
)

type customerController struct {
	customerService service.ICustomerService
}

func NewCustomerController(customerService service.ICustomerService) *customerController {
	return &customerController{customerService}
}

func (c *customerController) CreateCustomer(e echo.Context) error {
	var pagination utils.Pagination
	var input model.CustomerDto
	e.Bind(&input)
	customer := model.ConvertDtoToModel(input)
	pagination = input.Pagination

	var err error
	err, pagination.CurrentPage = c.customerService.SaveCustomer(&customer)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}

	return utils.NewSuccessPaginationResponse(e, customer, pagination)
}

func (c *customerController) GetAllCustomer(e echo.Context) error {
	var filter model.CustomerFilter
	if err := e.Bind(&filter); err != nil {
		return err
	}
	var pagination utils.Pagination
	if err := e.Bind(&pagination); err != nil {
		return err
	}

	res, err := c.customerService.GetAllCustomer(filter, &pagination)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}

	return utils.NewSuccessPaginationResponse(e, res, pagination)
}