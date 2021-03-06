package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/infrastructure/config"
)

var DB *gorm.DB

func init(){
	conn, err := gorm.Open(mysql.Open(config.Config.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect mysql server, error: %s", err.Error())
		return
	}
	DB = conn
}
