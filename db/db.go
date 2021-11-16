package db

import (
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ConnDB *gorm.DB

type Track struct {
	gorm.Model
	Artist string //`json:"Artist" gorm:"<-"`
	Name   string //`json:"Name" gorm:"<-"`
	Genre  string //`json:"Genre" gorm:"<-"`
}

func ConnectToDB() (*gorm.DB, error) {
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "TracksDB",
		ParseTime: true,
	}
	var err error
	ConnDB, err = gorm.Open("mysql", cfg.FormatDSN())
	return ConnDB, err
}

func CreateTrack(db *gorm.DB, track *Track) error {
	err := db.Create(track).Error
	return err
}

func SelectTrack(db *gorm.DB, track *Track, id int) error {
	err := db.First(track, id).Error
	return err
}

func SelectTracks(db *gorm.DB, tracks *[]Track) error {
	err := db.Find(tracks).Error
	return err
}

func DeleteTracks(db *gorm.DB, track *Track, id int) error {
	err := db.Where("id=?", id).Delete(track).Error
	return err
}

func UpdateTrack(db *gorm.DB, track *Track, id int) error {
	err := db.Where("id=?").Error
	if err != nil {
		return err
	}
	db.Save(track)
	return nil
}
