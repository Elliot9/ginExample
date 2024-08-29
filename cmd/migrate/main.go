package main

import (
	"fmt"
	"log"

	"github/elliot9/ginExample/config"
	"github/elliot9/ginExample/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.Load(".env")
	dbConfig := config.WDbSetting

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Can not connect to DB: %v", err)
	}

	// begin transaction
	tx := db.Begin()
	err = tx.AutoMigrate(
		&models.Admin{},
		&models.Article{},
		&models.User{},
		&models.UserRefreshToken{},
		// append other models here...
	)

	if err != nil {
		tx.Rollback()
		log.Fatalf("Migrate Fail: %v", err)
	}

	tx.Commit()
	fmt.Println("Migrate Done!")
}
