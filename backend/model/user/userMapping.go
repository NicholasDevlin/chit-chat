package user

import "github.com/jinzhu/gorm"

func ConvertReqToDto(input UserReq) *UserDto {
	return &UserDto{
		UUID:        input.UUID,
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}
}

func ConvertDtoToModel(input UserDto) *User {
	return &User{
		Model: gorm.Model{
			ID:        input.Id,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		},
		UUID:        input.UUID,
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}
}

func ConvertModelToDto(input User) *UserDto {
	return &UserDto{
		Id:          input.ID,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		UUID:        input.UUID,
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
	}
}

func ConvertDtoToRes(input UserDto) *UserRes {
	return &UserRes{
		UUID:        input.UUID,
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
	}
}
