package account_test

import (
	accountPackage "github.com/magleff/gobro/features/account"
	"github.com/stretchr/testify/mock"
)

type MockAccountDatastore struct {
	mock.Mock
}

func (m MockAccountDatastore) Create(account accountPackage.Account) {
	m.Called(account)
}

func (m MockAccountDatastore) List() []accountPackage.Account {
	m.Called()
	return nil
}

func (m MockAccountDatastore) Current() *accountPackage.Account {
	args := m.Called()
	if account := args.Get(0); account == nil {
		return nil
	} else {
		return account.(*accountPackage.Account)
	}
}
