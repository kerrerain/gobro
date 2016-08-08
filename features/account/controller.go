package account

type AccountController interface {
	Create(string) error
	List() []Account
	Current() *Account
}

type AccountControllerImpl struct {
	AccountDatastore AccountDatastore
}

func NewAccountController() AccountController {
	instance := new(AccountControllerImpl)
	instance.AccountDatastore = new(AccountDatastoreImpl)
	return instance
}

func (self AccountControllerImpl) Create(name string) error {
	account := Account{
		Name: name,
	}
	self.AccountDatastore.Create(account)
	return nil
}

func (self AccountControllerImpl) List() []Account {
	return self.AccountDatastore.List()
}

func (self AccountControllerImpl) Current() *Account {
	return self.AccountDatastore.Current()
}
