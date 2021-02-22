package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/mrrizal/rest-api-blog/configs"
	log "github.com/sirupsen/logrus"
)

var config configs.Config

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "", "migrate command")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.LoadConfig()

	app := FiberRouter().InitRouter()

	if migrate != "" {
		if err := postgresSQLHandler.Migrate(migrate, config.DBURL); err != nil {
			log.Fatal(err)
		}
	} else {
		app.Listen(":3000")
	}
}
