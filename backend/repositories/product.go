package repositories

import (
	// uuid "github.com/satori/go.uuid"
	"math"
	"myapp/backend/model"
	"myapp/backend/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IProductRpository interface {
	SaveProduct(input *model.Product) (error, int)
	GetAllProduct(filter model.ProductFilter, pagination *utils.Pagination) ([]model.Product, error)
	GetProduct(filter *model.ProductFilter) (model.Product, error)
	GetProductDetail(filter *model.ProductDetailFilter) (model.ProductDetail, error)
	DeleteProduct(id string) error
}

type ProductRepository struct {
	db *gorm.DB
}

// GetProductDetail implements IProductRpository.
func (c *ProductRepository) GetProductDetail(filter *model.ProductDetailFilter) (model.ProductDetail, error) {
	var detail model.ProductDetail
	query := c.db.Model(&model.ProductDetail{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	err := query.First(&detail).Error
	return detail, err
}

// DeleteProduct implements IProductRpository.
func (*ProductRepository) DeleteProduct(id string) error {
	panic("unimplemented")
}

// GetAllProduct implements IProductRpository.
func (c *ProductRepository) GetAllProduct(filter model.ProductFilter, pagination *utils.Pagination) ([]model.Product, error) {
	var Products []model.Product

	query := c.db.Model(&model.Product{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	if filter.SortName {
		query = query.Order("name DESC")
	}

	query.Preload("ProductDetail").Model(&model.Product{}).Count(&pagination.TotalRecords)
	pagination.AllPages = int64(math.Ceil(float64(pagination.TotalRecords) / float64(pagination.PageSize)))

	err := query.Offset((pagination.CurrentPage - 1) * pagination.PageSize).Limit(pagination.PageSize).Find(&Products).Error

	return Products, err
}

// GetProduct implements IProductRpository.
func (c *ProductRepository) GetProduct(filter *model.ProductFilter) (model.Product, error) {
	var Product model.Product
	query := c.db.Model(&model.Product{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	err := query.First(&Product).Error
	return Product, err
}

// SaveProduct implements IProductRpository.
func (c *ProductRepository) SaveProduct(input *model.Product) (error, int) {
	if input.UUID == uuid.Nil {
		input.UUID = uuid.NewV4()
	}

	err := c.db.Save(&input).Error

	var position int64
	c.db.Model(&model.Product{}).
		Where("name <= ?", input.Name).
		Count(&position)

	pageSize := 10
	pageNumber := int(position / int64(pageSize))
	if position%int64(pageSize) > 0 {
		pageNumber++
	}

	return err, pageNumber
}

func NewProductRepository(db *gorm.DB) IProductRpository {
	return &ProductRepository{db: db}
}
