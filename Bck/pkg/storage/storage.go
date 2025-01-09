package storage

import (
	"Golang/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dsn := "user=postgres password=admin dbname=dbtest port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения базы данных dbtest: %v", err)
	}

	err = db.AutoMigrate(&models.Advertisement{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	log.Println("База данных успешно создана и миграция завершена.")

	return db
}
