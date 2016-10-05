package account

import (
	"github.com/magleff/gobro/entities"
	"gopkg.in/mgo.v2/bson"
)

type AccountController interface {
	List(userId bson.ObjectId) ([]entities.Account, error)
	Create(userId bson.ObjectId, accountName string) error
}

type AccountControllerImpl struct{}
