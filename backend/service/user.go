package service

import (
	"myapp/backend/model/user"
	"myapp/backend/repositories"
	"myapp/backend/utils"
	"myapp/backend/utils/bcrypt"
	"myapp/backend/utils/errors"

	uuid "github.com/satori/go.uuid"
)

type IUserService interface {
	RegisterUser(input user.UserReq) (user.UserRes, error)
	LoginUser(input user.UserReq) (user.UserRes, error)
	GetAllUser(filter user.UserFilter, pagination *utils.Pagination) ([]user.UserRes, error)
	SaveUser(input user.UserReq) (user.UserRes, error)
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	userRepository repositories.IUserRepository
}

// DeleteUser implements IUserService.
func (u *userService) DeleteUser(id uuid.UUID) error {
	res, err := u.userRepository.GetUser(user.UserFilter{UUID: id})
	if err != nil && res.UUID == id {
		return errors.ERR_USER_NOT_FOUND
	}

	err = u.userRepository.DeleteUser(id.String())
	if err != nil {
		return errors.ERR_DELETE_DATA
	}
	return nil
}

// GetAllUser implements IUserService.
func (u *userService) GetAllUser(filter user.UserFilter, pagination *utils.Pagination) ([]user.UserRes, error) {
	if pagination.PageSize == 0 {
		pagination.PageSize = 10
	}
	if pagination.CurrentPage == 0 {
		pagination.CurrentPage = 1
	}
	res, err := u.userRepository.GetAllUser(filter, pagination)
	if err != nil {
		return nil, err
	}

	var resUser []user.UserRes
	for i := 0; i < len(res); i++ {
		resUser = append(resUser, *user.ConvertDtoToRes(res[i]))
	}

	return resUser, nil
}

// LoginUser implements IUserService.
func (u *userService) LoginUser(input user.UserReq) (user.UserRes, error) {
	if input.Email == "" {
		return user.UserRes{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if input.Password == "" {
		return user.UserRes{}, errors.ERR_PASSWORD_IS_EMPTY
	}

	res, err := u.userRepository.GetUser(user.UserFilter{Email: input.Email})
	if err != nil {
		return user.UserRes{}, errors.ERR_USER_NOT_FOUND
	}

	err = bcrypt.CheckPassword(input.Password, res.Password)
	if err != nil {
		return user.UserRes{}, errors.ERR_WRONG_PASSWORD
	}

	return *user.ConvertDtoToRes(res), nil
}

// RegisterUser implements IUserService.
func (u *userService) RegisterUser(input user.UserReq) (user.UserRes, error) {
	if input.PhoneNumber == "" {
		return user.UserRes{}, errors.ERR_PHONE_NUMBER_IS_EMPTY
	}
	if input.Email == "" {
		return user.UserRes{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if input.Password == "" {
		return user.UserRes{}, errors.ERR_PASSWORD_IS_EMPTY
	}
	if input.Name == "" {
		return user.UserRes{}, errors.ERR_NAME_IS_EMPTY
	}

	isExist, err := u.userRepository.GetUser(user.UserFilter{Email: input.Email})
	if err == nil && isExist.Email == input.Email {
		return user.UserRes{}, errors.ERR_EMAIL_IS_TAKEN
	}

	hashPass, err := bcrypt.HashPassword(input.Password)
	if err != nil {
		return user.UserRes{}, errors.ERR_BCRYPT_PASSWORD
	}

	input.Password = hashPass
	res, err := u.userRepository.RegisterUser(*user.ConvertReqToDto(input))
	if err != nil {
		return user.UserRes{}, err
	}
	return *user.ConvertDtoToRes(res), nil
}

// SaveUser implements IUserService.
func (u *userService) SaveUser(input user.UserReq) (user.UserRes, error) {
	res, err := u.userRepository.GetUser(user.UserFilter{UUID: input.UUID})
	if err != nil {
		return user.UserRes{}, errors.ERR_USER_NOT_FOUND
	}
	if input.PhoneNumber != "" {
		res.PhoneNumber = input.PhoneNumber
	}
	if input.Name != "" {
		res.Name = input.Name
	}
	if input.Email != "" {
		res.Email = input.Email
	}

	err = u.userRepository.SaveUser(&res)
	if err != nil {
		return user.UserRes{}, errors.ERR_SAVE_DATA
	}
	return *user.ConvertDtoToRes(res), nil
}

func NewUserService(repo repositories.IUserRepository) IUserService {
	return &userService{userRepository: repo}
}
