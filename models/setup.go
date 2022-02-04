package models

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func SetupDB(){
	log.Println("Connecting to the Database")
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, viper.GetString("database_user"), viper.GetString("database_password"), viper.GetString("database_host"), viper.GetInt("database_port"),viper.GetString("database_name"))
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to the database")
	}
	log.Println("Connected to the Database")
}
