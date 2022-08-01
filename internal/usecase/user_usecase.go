package usecase

import (
	"cleanArch/internal/domain/model"
	"cleanArch/internal/dto"
	"cleanArch/internal/usecase/presenter"
	"cleanArch/internal/usecase/repository"
)

type UserUseCase interface {
	Get() ([]*dto.User, error)
}

type userUseCase struct {
	repository.UserRepository
	presenter.UserPresenter
}

func NewUserUseCase(
	userRepository repository.UserRepository,
	userPresenter presenter.UserPresenter,
) UserUseCase {
	return &userUseCase{
		UserRepository: userRepository,
		UserPresenter:  userPresenter,
	}
}

func (uc *userUseCase) Get() ([]*dto.User, error) {
	var users []*model.User

	users, err := uc.UserRepository.FindAll(users)
	if err != nil {
		return nil, err
	}

	return uc.UserPresenter.ResponseUsers(users), nil
}
