package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	DbAddr string
	AppEnv string
	Port   int
	JwtKey string
}

func New() *Config {
	currConfig := new(Config)
	currConfig.load()

	return currConfig
}

func (c *Config) load() {
	c.AppEnv = os.Getenv("APP_ENV")
	c.DbAddr = os.Getenv("DB_ADDR")

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Unable to get port due: %s", err)
	}

	c.Port = port

	c.JwtKey = os.Getenv("JWT_KEY")
}
