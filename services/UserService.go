package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/mrrizal/rest-api-blog/interfaces"
	"github.com/mrrizal/rest-api-blog/models"
)

type UserService struct {
	interfaces.IUserRepository
}

var validate *validator.Validate

func (service *UserService) SaveUserService(user models.UserModel) (models.UserModel, error) {
	validate = validator.New()
	if err := validate.Struct(user); err != nil {
		return models.UserModel{}, err
	}

	user, err := service.SaveUser(user)
	if err != nil {
		return models.UserModel{}, err
	}
	return user, nil
}
