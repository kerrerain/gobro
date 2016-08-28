package user

type Controller interface {
	OpenAccount(userName string, accountName string) error
	Create(userName string) error
	InitDefault()
}

type Impl struct{}
