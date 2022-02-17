package database

import (
	"fmt"
	"v1/admin"
	"v1/config"
	"v1/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {
	var (
		host     = config.Env("DB_HOST")
		port     = config.Env("DB_PORT")
		username = config.Env("DB_USER")
		password = config.Env("DB_PASSWORD")
		dbname   = config.Env("DB_NAME")
	)
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var user = user.User{}
	var admin = admin.Admin{}
	db.AutoMigrate(&user, &admin)
	return db, nil
}
