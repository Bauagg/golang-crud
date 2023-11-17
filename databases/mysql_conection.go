package databases

import (
	"belajar-api-goleng/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	var err error
	DB, err = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})

	if err != nil {
		panic("connect databases MYSQL gagal")
	}

	log.Println("Connection databases success")
}

// func CloseMysql() {
// 	if DB != nil {
// 		sqlDb, _ := DB.DB()
// 		sqlDb.Close()
// 	}
// }
