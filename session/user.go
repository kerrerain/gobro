package session

import (
	"github.com/magleff/gobro/common"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
	"log"
	"sync"
)

var user *entities.User
var userSession UserSession
var once sync.Once

type UserSession interface {
	GetUser() *entities.User
}

type UserSessionImpl struct{}

func (u UserSessionImpl) GetUser() *entities.User {
	return GetUserDo(dao.UserDaoImpl{})
}

// This is default behavior until there is a real user management.
// For the moment, it is only to make sure that a user ID is available.
func GetUserDo(userDao dao.UserDao) *entities.User {
	once.Do(func() {
		defaultUser, err := userDao.FindByName(common.DEFAULT_USER_NAME)
		// Until there is a real user management and a real logger,
		// consider this as a dirty placeholder in case of an error
		if err != nil {
			log.Fatal("Impossible to fetch the default user.")
		}
		user = defaultUser
	})
	return user
}

func InitUserSession() {
	userSession = UserSessionImpl{}
}

func GetSession() UserSession {
	return userSession
}
