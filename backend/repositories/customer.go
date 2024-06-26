package repositories

import (
	// uuid "github.com/satori/go.uuid"
	"math"
	"myapp/backend/model"
	"myapp/backend/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ICustomerRpository interface {
	SaveCustomer(input *model.Customer) (error, int)
	GetAllCustomer(filter model.CustomerFilter, pagination *utils.Pagination) ([]model.Customer, error)
	GetCustomer(filter *model.CustomerFilter) (model.Customer, error)
	DeleteCustomer(id string) error
}

type customerRepository struct {
	db *gorm.DB
}

// DeleteCustomer implements ICustomerRpository.
func (*customerRepository) DeleteCustomer(id string) error {
	panic("unimplemented")
}

// GetAllCustomer implements ICustomerRpository.
func (c *customerRepository) GetAllCustomer(filter model.CustomerFilter, pagination *utils.Pagination) ([]model.Customer, error) {
	var customers []model.Customer

	query := c.db.Model(&model.Customer{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.SortName {
		query = query.Order("name ASC")
	}

	query.Model(&model.Customer{}).Count(&pagination.TotalRecords)
	pagination.AllPages = int64(math.Ceil(float64(pagination.TotalRecords) / float64(pagination.PageSize)))

	err := query.Offset((pagination.CurrentPage - 1) * pagination.PageSize).Limit(pagination.PageSize).Find(&customers).Error

	return customers, err
}

// GetCustomer implements ICustomerRpository.
func (c *customerRepository) GetCustomer(filter *model.CustomerFilter) (model.Customer, error) {
	var customer model.Customer
	query := c.db.Model(&model.Customer{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	err := query.First(&customer).Error
	return customer, err
}

// SaveCustomer implements ICustomerRpository.
func (c *customerRepository) SaveCustomer(input *model.Customer) (error, int) {
	if input.UUID == uuid.Nil {
		input.UUID = uuid.NewV4()
	}

	err := c.db.Save(&input).Error

	var position int64
	c.db.Model(&model.Customer{}).
		Where("name <= ?", input.Name).
		Count(&position)

	pageSize := 10
	pageNumber := int(position / int64(pageSize))
	if position%int64(pageSize) > 0 {
		pageNumber++
	}

	return err, pageNumber
}

func NewCustomerRepository(db *gorm.DB) ICustomerRpository {
	return &customerRepository{db: db}
}
