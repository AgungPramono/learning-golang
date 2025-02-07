package repository

import (
	"github.com/stretchr/testify/mock"
	"golang-unit-test/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock // mocking repository
}

func (repository CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
