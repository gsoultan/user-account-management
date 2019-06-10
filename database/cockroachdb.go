package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type cockroachdb struct {
	db *gorm.DB
}

func NewCockroachDB(connectionString string) (database Database, rerr error) {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error: Connection failed to database:", err.Error())
	}
	return &cockroachdb{db: db}, nil
}

func (c *cockroachdb) GetConnection() interface{} {
	return c.db
}
