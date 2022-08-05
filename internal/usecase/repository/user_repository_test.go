package repository

import (
	"fmt"
	"testing"

	"cleanArch/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedDB struct {
	mock.Mock
}

func newMockedDB(result []*model.User, err error) *mockedDB {
	m := &mockedDB{}

	m.
		On("FindAll").
		Return(result, err)

	return m
}

func (m *mockedDB) FindAll() ([]*model.User, error) {
	arguments := m.Called()

	return arguments.Get(0).([]*model.User), arguments.Error(1)
}

func (m *mockedDB) Close() error {
	return nil
}

func TestDatabaseCallIsCalledOneTime(t *testing.T) {
	mockedDb := newMockedDB([]*model.User{}, nil)

	repository := NewUserRepository(mockedDb)

	repository.FindAll()

	mockedDb.AssertNumberOfCalls(t, "FindAll", 1)
}

func TestSuccessfullyGetListOfUsersFromDB(t *testing.T) {
	mockedDb := newMockedDB([]*model.User{}, nil)

	repository := NewUserRepository(mockedDb)

	received, _ := repository.FindAll()

	assert.Equal(t, []*model.User{}, received)
}

func TestPropagateErrorFromDB(t *testing.T) {
	expectedError := fmt.Errorf("I want this Error")
	mockedDb := newMockedDB([]*model.User{}, expectedError)

	repository := NewUserRepository(mockedDb)

	received, err := repository.FindAll()

	assert.Nil(t, received)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}
