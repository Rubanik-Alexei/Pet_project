package db

import (
	//"fmt"

	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ConnDB *gorm.DB

type Track struct {
	gorm.Model
	id     int
	artist string
	name   string
	genre  string
}

func ConnectToDB() (*gorm.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "StationsDB",
	}
	var err error
	ConnDB, err = gorm.Open("mysql", cfg.FormatDSN())
	return ConnDB, err
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// fmt.Println("Connection Opened to Database")
}
