package account

import (
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

type AccountController interface {
	List() []entities.Account
	Create(userId bson.ObjectId, accountName string) error
}

type AccountControllerImpl struct{}
