package restrict

import (
	"integration-test/lib"
	"integration-test/models"
)

func ChecklUserExisting(email string, user *models.User) bool {
	err := lib.FindUserByEmail(email, user)
	if err == nil {
		return true
	}
	return false
}

// func TestChecklUserExisting(email string, user *models.TestUser) bool {
// 	err := lib.TestFindUserByEmail(email, user)
// 	if err == nil {
// 		return true
// 	}
// 	return false
// }
