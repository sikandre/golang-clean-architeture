package controller

import (
	"net/http"

	"cleanArch/internal/usecase"
)

type UserController interface {
	GetUsers(ctx Context) error
}

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(userUseCase usecase.UserUseCase) UserController {
	return &userController{userUseCase}
}

type Binder struct {
	Id string `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id"`
}

func (uc *userController) GetUsers(ctx Context) error {
	dto, err := uc.userUseCase.Get()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, dto)
}
