package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spinales/quiz-api/api"
	"github.com/spinales/quiz-api/models"
	storage "github.com/spinales/quiz-api/storage/postgres"
	"github.com/spinales/quiz-api/util"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	local bool
	fake  bool
)

func init() {
	flag.BoolVar(&local, "local", false, "Run the server in local mode, run without database server using sqlite.")
	flag.BoolVar(&fake, "fake", false, "Add fake data on dabatase.")
	flag.Parse()
}

func main() {
	var db *gorm.DB

	config, err := util.NewConfig()
	if err != nil {
		log.Fatalln("Cannot load config file(.env): ", err)
	}

	if local {
		db, err = gorm.Open(sqlite.Open(":memory:"))
	} else {
		dsc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Host, config.User, config.Password, config.DBName, config.Port)
		db, err = gorm.Open(postgres.Open(dsc), &gorm.Config{})
		db.Exec(`CREATE TYPE role AS ENUM ('admin','player');`)
	}
	if err != nil {
		log.Fatalln("Cannot connect to database: ", err)
	}

	db.AutoMigrate(&models.User{}, &models.Question{}, &models.Answer{})

	if fake {
		storage.FakeData(db)
	}

	server, err := api.NewServer(db, &config)
	if err != nil {
		log.Fatalln("Cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("Cannot start server: ", err)
	}
}
