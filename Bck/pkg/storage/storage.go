package storage

import (
	"Golang/pkg/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения базы данных dbtest: %v", err)
	}

	err = db.AutoMigrate(&models.Advertisement{}, &models.Photo{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	log.Println("База данных успешно создана и миграция завершена.")

	return db
}

const (
	host     = "localhost"
	user     = "postgres"
	password = "admin"
	dbname   = "dbtest"
	port     = 5432
)
