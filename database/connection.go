package database

import (
	"jwt/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn string = "root:@tcp(127.0.0.1:3306)/jwt_auth?charset=utf8mb4&parseTime=True&loc=Local"
// DB connection clone
var DB *gorm.DB
// Connection of Database
func Connection() {
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error Database")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}