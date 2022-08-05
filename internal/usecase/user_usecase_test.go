package usecase

import (
	"fmt"
	"testing"
	"time"

	"cleanArch/internal/domain/model"
	"cleanArch/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
	result []*model.User
	error  error
}

func newMockUserRepository(result []*model.User, err error) *mockUserRepository {
	m := &mockUserRepository{result: result, error: err}

	m.
		On("FindAll", mock.Anything).
		Return(result, err)

	return m
}

func (m *mockUserRepository) FindAll() ([]*model.User, error) {
	arguments := m.Called()

	return arguments.Get(0).([]*model.User), arguments.Error(1)
}

type mockUserPresenter struct {
	mock.Mock
	result []*dto.User
}

func newMockUserPresenter(result []*dto.User) *mockUserPresenter {
	m := &mockUserPresenter{result: result}

	m.
		On("ResponseUsers", mock.Anything).
		Return(result)

	return m
}

func (m *mockUserPresenter) ResponseUsers(users []*model.User) []*dto.User {
	arguments := m.Called(users)

	return arguments.Get(0).([]*dto.User)
}

func TestAssertDependenciesAreCalledOneTime(t *testing.T) {
	mockedRepositoryResult, mockedPresenterResult := successMockedResults()

	mockedRepository := newMockUserRepository(mockedRepositoryResult, nil)
	mockedPresenter := newMockUserPresenter(mockedPresenterResult)

	useCase := NewUserUseCase(mockedRepository, mockedPresenter)

	useCase.Get()

	mockedRepository.AssertNumberOfCalls(t, "FindAll", 1)
	mockedPresenter.AssertNumberOfCalls(t, "ResponseUsers", 1)
}

func TestGetAllUsersFromRepository(t *testing.T) {
	mockedRepositoryResult, mockedPresenterResult := successMockedResults()

	mockedRepository := newMockUserRepository(mockedRepositoryResult, nil)
	mockedPresenter := newMockUserPresenter(mockedPresenterResult)

	useCase := NewUserUseCase(mockedRepository, mockedPresenter)

	received, _ := useCase.Get()

	mockedRepository.AssertExpectations(t)
	mockedPresenter.AssertExpectations(t)

	assert.Equal(t, mockedPresenter.result, received)
}

func TestGetAllUsersFromRepositoryWhenNoUsersExist(t *testing.T) {
	mockedRepositoryResult, mockedPresenterResult := noValuesMockedResults()

	mockedRepository := newMockUserRepository(mockedRepositoryResult, fmt.Errorf("dont care"))
	mockedPresenter := newMockUserPresenter(mockedPresenterResult)

	useCase := NewUserUseCase(mockedRepository, mockedPresenter)

	received, _ := useCase.Get()

	mockedRepository.AssertExpectations(t)
	mockedPresenter.AssertNumberOfCalls(t, "ResponseUsers", 0)

	assert.Nil(t, received)
}

func successMockedResults() ([]*model.User, []*dto.User) {
	mockedRepositoryResult := []*model.User{
		{
			Id:        1,
			Name:      "Bob",
			Age:       "30",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}

	mockedPresenterResult := []*dto.User{
		{
			Name: "Bob",
			Age:  "30",
		},
	}

	return mockedRepositoryResult, mockedPresenterResult
}

func noValuesMockedResults() ([]*model.User, []*dto.User) {
	return []*model.User{}, []*dto.User{}
}
