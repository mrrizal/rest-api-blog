package infrastructures

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/mrrizal/rest-api-blog/models"

	// go migrate postgres driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// go migrate file
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresSQLHandler struct {
	Conn *gorm.DB
}

func (handler *PostgresSQLHandler) Connect(dsn string) error {
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if connect fail
	if err != nil {
		r, _ := regexp.Compile("dbname=([a-zA-Z]+)")
		dbName := r.FindStringSubmatch(dsn)[1]
		dsn = strings.Replace(dsn, dbName, "postgres", 1)

		// connect to postgres db then create database
		dbConn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		handler.Conn = dbConn
		handler.Execute(fmt.Sprintf("CREATE DATABASE %s", dbName))

		// try connect to db again
		dsn = strings.Replace(dsn, "dbname=postgres", fmt.Sprintf("dbname=%s", dbName), 1)
		dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		handler.Conn = dbConn
	}
	handler.Conn = dbConn
	return nil
}

func (handler *PostgresSQLHandler) Execute(statement string) {
	handler.Conn.Exec(statement)

}

func (handler *PostgresSQLHandler) Migrate(command, dbURL string) error {
	log.Info(fmt.Sprintf("run migration %s", command))
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		return err
	}

	switch command {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "drop":
		err = m.Drop()
	default:
		log.Info("migration command not avaible")
		return nil
	}

	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("run migration %s done", command))
	return nil
}

func (handler *PostgresSQLHandler) CreateUser(user *models.UserModel, tableName string) error {
	result := handler.Conn.Table(tableName).Create(user)
	return result.Error
}
