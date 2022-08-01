package presenter

import (
	"cleanArch/internal/domain/model"
	"cleanArch/internal/dto"
)

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*dto.User
}

type userPresenter struct{}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(users []*model.User) []*dto.User {
	usersDto := make([]*dto.User, len(users))

	for i, user := range users {
		usersDto[i] = &dto.User{
			Name: user.Name,
			Age:  user.Age,
		}
	}

	return usersDto
}
