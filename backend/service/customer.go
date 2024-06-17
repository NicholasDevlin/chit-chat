package service

import (
	"myapp/backend/model"
	"myapp/backend/repositories"
	"myapp/backend/utils"
	"myapp/backend/utils/errors"

	uuid "github.com/satori/go.uuid"
)

type ICustomerService interface {
	GetAllCustomer(filter model.CustomerFilter, pagination *utils.Pagination) ([]model.Customer, error)
	GetCustomer(filter model.CustomerFilter) (model.Customer, error)
	SaveCustomer(input *model.Customer) (error, int)
	// DeleteTransaction(id uuid.UUID) (transaction.TransactionRes, error)
}

type customerService struct {
	customerRepository repositories.ICustomerRpository
}

// GetAllCustomer implements ICustomerService.
func (c *customerService) GetAllCustomer(filter model.CustomerFilter, pagination *utils.Pagination) ([]model.Customer, error) {
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}
	if pagination.CurrentPage == 0 {
		pagination.CurrentPage = 1
	}
	res, err := c.customerRepository.GetAllCustomer(filter, pagination)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetCustomer implements ICustomerService.
func (*customerService) GetCustomer(filter model.CustomerFilter) (model.Customer, error) {
	panic("unimplemented")
}

// SaveCustomer implements ICustomerService.
func (c *customerService) SaveCustomer(input *model.Customer) (error, int) {
	var err error
	if input.Address == "" {
		return errors.ERR_CREATE_CUSTOMER, 1
	}
	if input.Age == 0 {
		return errors.ERR_AGE_IS_EMPTY, 1
	}
	if input.Name == "" {
		return errors.ERR_NAME_IS_EMPTY, 1
	}
	var newData model.Customer
	newData = *input
	if input.UUID != uuid.Nil {
		*input, err = c.customerRepository.GetCustomer(&model.CustomerFilter{UUID: newData.UUID})
		if err != nil {
			return errors.ERR_CREATE_CUSTOMER,1 
		}
		input.Address = newData.Address
		input.Age = newData.Age
		input.Name = newData.Name
	}

	err, pageNumber := c.customerRepository.SaveCustomer(input)
	if err != nil {
		return errors.ERR_CREATE_CUSTOMER, 1
	}
	return err, pageNumber
}

func NewCustomerService(repo repositories.ICustomerRpository) ICustomerService {
	return &customerService{
		customerRepository: repo,
	}
}
