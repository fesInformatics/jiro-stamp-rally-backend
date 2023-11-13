package dbconfig

import (
	"os"
	"strconv"
)

type DBconfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

func NewDBConfig() *DBconfig {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &DBconfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		DBName:   os.Getenv("DB_DATABASE"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}
