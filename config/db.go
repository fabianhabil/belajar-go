package config

import (
	"os"
)

type DB struct {
	Host     string
	Port     string
	SslMode  string
	Name     string
	User     string
	Password string
}

var db = &DB{}

func DBCfg() *DB {
	return db
}

func LoadDBCfg() {
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Name = os.Getenv("DB_NAME")
	db.SslMode = os.Getenv("DB_SSLMODE")
}
