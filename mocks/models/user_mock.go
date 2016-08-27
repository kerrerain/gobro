package models

import (
	"github.com/magleff/gobro/models"
	"github.com/stretchr/testify/mock"
)

type User struct {
	mock.Mock
}

func (m User) FindByName(userName string) *models.User {
	args := m.Called(userName)
	if user := args.Get(0); user == nil {
		return nil
	} else {
		return user.(*models.User)
	}
}

func (m User) Update(user models.User) {
	m.Called(user)
}
