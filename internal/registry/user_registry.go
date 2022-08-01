package registry

import (
	"cleanArch/internal/controller"
	"cleanArch/internal/usecase"
	"cleanArch/internal/usecase/presenter"
	"cleanArch/internal/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserUseCase())
}

func (r *registry) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() repository.UserRepository {
	return repository.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() presenter.UserPresenter {
	return presenter.NewUserPresenter()
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		User: r.NewUserController(),
	}
}
