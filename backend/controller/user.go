package controller

import (
	"myapp/backend/model/user"
	"myapp/backend/service"
	"myapp/backend/utils"
	"myapp/backend/utils/middleware"
	"myapp/backend/utils/errors"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type userController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *userController {
	return &userController{userService}
}

func (u *userController) RegisterUsers(e echo.Context) error {
	var input user.UserReq
	e.Bind(&input)

	res, err := u.userService.RegisterUser(input)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}

	token, err := middleware.CreateToken(res.UUID, res.Name)
	if err != nil {
		return utils.NewErrorResponse(e, errors.ERR_TOKEN)
	}
	res.Token = token

	return utils.NewSuccessResponse(e, res)
}

func (u *userController) LoginUser(e echo.Context) error {
	var input user.UserReq
	e.Bind(&input)

	res, err := u.userService.LoginUser(input)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}

	token, err := middleware.CreateToken(res.UUID, res.Name)
	if err != nil {
		return utils.NewErrorResponse(e, errors.ERR_TOKEN)
	}
	res.Token = token

	return utils.NewSuccessResponse(e, res)
}

func (u *userController) GetAllUser(e echo.Context) error {
	var filter user.UserFilter
	if err := e.Bind(&filter); err != nil {
		return err
	}
	var pagination utils.Pagination
	if err := e.Bind(&pagination); err != nil {
		return err
	}

	res, err := u.userService.GetAllUser(filter, &pagination)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}

	return utils.NewSuccessPaginationResponse(e, res, pagination)
}

func (u *userController) SaveUser(e echo.Context) error {
	userId, err := middleware.ExtractToken(e)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}

	var input user.UserReq
	e.Bind(&input)
	uuid, err := uuid.FromString(e.Param("id"))
	if userId != uuid {
		return utils.NewErrorResponse(e, errors.ERR_UNAUTHORIZE)
	}
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}
	input.UUID = uuid
	res, err := u.userService.SaveUser(input)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}
	return utils.NewSuccessResponse(e, res)
}

func (u *userController) DeleteUser(e echo.Context) error {
	userId, err := middleware.ExtractToken(e)

	uuid, err := uuid.FromString(e.Param("id"))
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}

	if userId != uuid {
		return utils.NewErrorResponse(e, errors.ERR_UNAUTHORIZE)
	}

	err = u.userService.DeleteUser(uuid)
	if err != nil {
		return utils.NewErrorResponse(e, err)
	}
	return utils.NewSuccessResponse(e, user.UserRes{})
}