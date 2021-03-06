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
	"gorm.io/gorm"
)

var (
	local    bool
	fake     bool
	reset    bool
	cleanAll bool
)

func init() {
	flag.BoolVar(&local, "local", false, "Run the server in local mode, run without database server using sqlite.")
	flag.BoolVar(&fake, "fake", false, "Add fake data on dabatase.")
	flag.BoolVar(&reset, "reset", false, "Reset database to default data.")
	flag.BoolVar(&cleanAll, "new", false, "Erase all the data, without admin user.")
	flag.Parse()
}

func main() {
	var db *gorm.DB

	config, err := util.NewConfig()
	if err != nil {
		log.Fatalln("Cannot load config file(.env): ", err)
	}

	if local {
		// db, err = gorm.Open(sqlite.Open(":memory:"))
	} else {
		dsc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
			config.Host, config.User, config.Password, config.DBName, config.Port)
		db, err = gorm.Open(postgres.Open(dsc), &gorm.Config{})
		db.Exec(`CREATE TYPE role AS ENUM ('admin','player');`)
	}
	if err != nil {
		log.Fatalln("Cannot connect to database: ", err)
	}

	db.AutoMigrate(&models.User{}, &models.Question{})

	if fake {
		storage.FakeData(db)
	}

	if reset {
		storage.EraseData(db)
	}

	if cleanAll {
		storage.EraseData(db)
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
