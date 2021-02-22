package interfaces

import "github.com/mrrizal/rest-api-blog/models"

type IUserRepository interface {
	SaveUser(user models.UserModel) (models.UserModel, error)
	GetUsers(userID int) ([]models.UserModel, error)
	GetUser(username string) (models.UserModel, error)
}
