package repositories

import (
	"github.com/mrrizal/rest-api-blog/interfaces"
	"github.com/mrrizal/rest-api-blog/models"
)

type UserRepository struct {
	interfaces.IDBHandler
}

func (repository *UserRepository) SaveUser(user models.UserModel) (models.UserModel, error) {
	if err := repository.CreateUser(&user, user.TableName()); err != nil {
		return models.UserModel{}, err
	}
	return user, nil
}

func (repository *UserRepository) GetUsers(userID int) ([]models.UserModel, error) {
	return []models.UserModel{}, nil
}

func (repository *UserRepository) GetUser(username string) (models.UserModel, error) {
	return models.UserModel{}, nil
}
