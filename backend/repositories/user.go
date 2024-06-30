package repositories

import (
	"myapp/backend/model/user"
	"myapp/backend/utils"
	"math"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(data user.UserDto) (user.UserDto, error)
	GetAllUser(filter user.UserFilter, pagination *utils.Pagination) ([]user.UserDto, error)
	GetUser(filter user.UserFilter) (user.UserDto, error)
	SaveUser(input *user.UserDto) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

// DeleteUser implements IUserRepository.
func (u *userRepository) DeleteUser(id string) error {
	userData := user.User{}

	err := u.db.Delete(&userData, "uuid = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

// GetAllUser implements IUserRepository.
func (u *userRepository) GetAllUser(filter user.UserFilter, pagination *utils.Pagination) ([]user.UserDto, error) {
	var users []user.User

	query := u.db.Model(&user.User{})
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	if filter.SortByName {
		query = query.Order("name ASC")
	}

	query.Model(&user.User{}).Count(&pagination.TotalRecords)
	pagination.AllPages = int64(math.Ceil(float64(pagination.TotalRecords) / float64(pagination.PageSize)))

	err := query.Offset((pagination.CurrentPage - 1) * pagination.PageSize).Limit(pagination.PageSize).Find(&users).Error

	var resUser []user.UserDto
	for i := 0; i < len(users); i++ {
		resUser = append(resUser, *user.ConvertModelToDto(users[i]))
	}

	return resUser, err
}

// GetUser implements IUserRepository.
func (u *userRepository) GetUser(filter user.UserFilter) (user.UserDto, error) {
	var userData user.User
	query := u.db
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.UUID != uuid.Nil {
		query = query.Where("uuid = ?", filter.UUID)
	}
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}

	err := query.First(&userData).Error
	if err != nil {
		return user.UserDto{}, err
	}
	return *user.ConvertModelToDto(userData), nil
}

// SaveUser implements IUserRepository.
func (u *userRepository) SaveUser(input *user.UserDto) error {
	data := user.ConvertDtoToModel(*input)
	err := u.db.Save(&data).Error
	input = user.ConvertModelToDto(*data) 
	return err
}

func NewUsersRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (u *userRepository) RegisterUser(data user.UserDto) (user.UserDto, error) {
	dataUser := user.ConvertDtoToModel(data)
	dataUser.UUID = uuid.NewV4()
	err := u.db.Create(&dataUser).Error
	if err != nil {
		return user.UserDto{}, err
	}
	return *user.ConvertModelToDto(*dataUser), nil
}
