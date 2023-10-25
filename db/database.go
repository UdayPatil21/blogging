package db

import (
	"log"

	model "github.com/UdayPatil21/blogging/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Init() {
	//Init database

	MySqlConnection()
}

var DBInstance *gorm.DB
var DbError error

const (
	BlogDB = "Blogs"
)

func MySqlConnection() {
	dburl := "root:Digi@2023@/Blogging?parseTime=true"
	DBInstance, DbError = gorm.Open("mysql", dburl)
	if DbError != nil {
		log.Println("Cannot Connect to Database", DbError)
	}
	log.Println("Connected to Database")
	DBInstance.AutoMigrate(&model.BlogPost{})
}
