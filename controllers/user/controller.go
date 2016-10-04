package user

type UserController interface {
	Create(userName string) error
	InitDefault() error
}

type UserControllerImpl struct{}
