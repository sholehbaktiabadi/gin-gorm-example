package database

import (
	"v1/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=root dbname=go_prisma port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var user = user.User{}
	db.AutoMigrate(&user)
	return db, nil
}
