package util

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	User                string
	Password            string
	Port                string
	Host                string
	DBName              string
	ServerAddress       string
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
}

func NewConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	address := os.Getenv("Server_Address")
	key := os.Getenv("Token_Symmetric_Key")
	accessTokenDuration := os.Getenv("Access_Token_Duration")

	num, err := strconv.Atoi(accessTokenDuration)
	if err != nil {
		return Config{}, err
	}

	config := Config{user, password, port, host, name, address, key, time.Hour * time.Duration(num)}

	return config, nil
}
