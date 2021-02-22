package interfaces

import "github.com/mrrizal/rest-api-blog/models"

type IUserService interface {
	SaveUserService(user models.UserModel) (models.UserModel, error)
}
