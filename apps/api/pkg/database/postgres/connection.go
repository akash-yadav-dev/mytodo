package postgres

import (
	"fmt"
)

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Config struct {
	Connection Connection
}

func NewConnection(host, port, user, password, dbName string) Connection {
	return Connection{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}
}

func (c Connection) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
	)
}
