package session

import (
	"github.com/magleff/gobro/common"
	"github.com/magleff/gobro/models"
	"log"
	"sync"
)

var sessionUser *models.User
var once sync.Once

func GetCurrentUser() *models.User {
	return GetCurrentUserDo(models.User{})
}

// This is default behavior until there is a real user management.
// For the moment, it is only to make sure that a user ID is available.
func GetCurrentUserDo(entity models.UserEntity) *models.User {
	once.Do(func() {
		user, err := entity.FindByName(common.DEFAULT_USER_NAME)
		// Until there is a real user management and a real logger,
		// consider this as a dirty placeholder in case of an error
		if err != nil {
			log.Fatal("Impossible to fetch the default user.")
		}
		sessionUser = user
	})
	return sessionUser
}
