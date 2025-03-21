package db

import (
	"log"

	"git.sindadsec.ir/asm/backend/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Organization{}); err != nil {
		log.Panicln(err.Error())
	}

	if err := db.AutoMigrate(&models.Target{}); err != nil {
		log.Panicln(err.Error())
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Panicln(err.Error())
	}

	if err := db.AutoMigrate(&models.Notification{}); err != nil {
		log.Panicln(err.Error())
	}

	if err := db.AutoMigrate(&models.Permission{}); err != nil {
		log.Panicln(err.Error())
	}
}
