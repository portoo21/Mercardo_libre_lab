package repositories

import (
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

type ConnectionType int8

const (
	REGULAR   = 0
	SENSITIVE = 1
)

type DBUser struct {
	name     string
	password string
}

func getDbUser(conType ConnectionType) DBUser {
	switch conType {
	case REGULAR:
		return DBUser{name: os.Getenv("DB_USER"), password: os.Getenv("DB_PASSWORD")}
	case SENSITIVE:
		return DBUser{name: os.Getenv("DB_SENSITIVE_USER"), password: os.Getenv("DB_SENSITIVE_PASSWORD")}
	}
	return DBUser{}
}

func connectDb(user DBUser) (*pgx.Conn, error) {
	dbPort, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 0)
	if err != nil {
		return nil, err
	}

	conConfig := pgx.ConnConfig{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		User:     user.name,
		Password: user.password,
		Port:     uint16(dbPort),
	}
	return pgx.Connect(conConfig)
}
