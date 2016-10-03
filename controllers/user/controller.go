package user

type UserController interface {
	OpenAccount(userName string, accountName string) error
	Create(userName string) error
	InitDefault()
}

type UserControllerImpl struct{}
