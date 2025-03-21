package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(addr string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		log.Panicln(err.Error())
	}

	migrate(db)
	return db
}
