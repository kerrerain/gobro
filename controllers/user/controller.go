package user

type Controller interface {
	Open(userName string, accountName string) error
}

type Impl struct{}
