package driver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresDB struct {
	DB *gorm.DB
}

var Postgres = &PostgresDB{}

func Connect(dbHost string, port string, username string, password string, dbName string) *PostgresDB {
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, port, username, dbName, password)
	db, err := gorm.Open("postgres", dbUri)
	if err !=nil {
		panic(err)
	}
	Postgres.DB = db
	return Postgres
}
