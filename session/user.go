package session

import (
	"github.com/magleff/gobro/common"
	"github.com/magleff/gobro/dao"
	"github.com/magleff/gobro/entities"
	"log"
	"sync"
)

var sessionUser *entities.User
var once sync.Once

func GetCurrentUser() *entities.User {
	return GetCurrentUserDo(dao.UserDaoImpl{})
}

// This is default behavior until there is a real user management.
// For the moment, it is only to make sure that a user ID is available.
func GetCurrentUserDo(userDao dao.UserDao) *entities.User {
	once.Do(func() {
		user, err := userDao.FindByName(common.DEFAULT_USER_NAME)
		// Until there is a real user management and a real logger,
		// consider this as a dirty placeholder in case of an error
		if err != nil {
			log.Fatal("Impossible to fetch the default user.")
		}
		sessionUser = user
	})
	return sessionUser
}
