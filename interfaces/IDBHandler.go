package interfaces

import "github.com/mrrizal/rest-api-blog/models"

type IDBHandler interface {
	Execute(statement string)
	CreateUser(user *models.UserModel, tableName string) error
}
