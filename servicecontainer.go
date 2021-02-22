package main

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/mrrizal/rest-api-blog/controllers"
	"github.com/mrrizal/rest-api-blog/infrastructures"
	"github.com/mrrizal/rest-api-blog/repositories"
	"github.com/mrrizal/rest-api-blog/services"
)

type IServiceContainer interface {
	InjectUserController() controllers.UserController
}

var postgresSQLHandler infrastructures.PostgresSQLHandler

type kernel struct{}

func parseDbURL() (string, error) {
	u, err := url.Parse(config.DBURL)
	if err != nil {
		return "", err
	}

	host, port, _ := net.SplitHostPort(u.Host)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s %s", host, u.User.Username(),
		strings.Replace(u.Path, "/", "", 1), port, u.RawQuery)
	return dsn, nil
}

func (k *kernel) InjectUserController() controllers.UserController {
	// parse db url
	dsn, err := parseDbURL()
	if err != nil {
		log.Fatal(err)
	}

	if err := postgresSQLHandler.Connect(dsn); err != nil {
		log.Fatal(err)
	}

	userRepository := &repositories.UserRepository{&postgresSQLHandler}
	userService := &services.UserService{userRepository}
	userController := controllers.UserController{userService}
	return userController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
