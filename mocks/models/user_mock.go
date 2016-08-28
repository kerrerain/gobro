package models

import (
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/mock"
)

type User struct {
	mock.Mock
}

func (m User) FindByName(userName string) (*models.User, error) {
	args := m.Called(userName)
	if user := args.Get(0); user == nil {
		return nil, args.Error(1)
	} else {
		return user.(*models.User), args.Error(1)
	}
}

func (m User) Update(user models.User) {
	m.Called(user)
}

func (m User) Create(user models.User) {
	m.Called(user)
}

func (m User) UpdateAccount(user models.User, account models.Account) {
	m.Called(user, account)
}
