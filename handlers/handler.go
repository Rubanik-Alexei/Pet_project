package handlers

import (
	"Pet_project/db"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type DB struct {
	Db *gorm.DB
}

func NewDB() *DB {
	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to DB")
	dbConn.AutoMigrate(&db.Track{})
	return &DB{Db: dbConn}

}

func (Db *DB) GetAlltracks(ctx *fiber.Ctx) error {
	tracks := []db.Track{}
	err := db.SelectTracks(Db.Db, &tracks)
	if err != nil {
		ctx.SendString("Cannot retrieve tracks from db")
		return err
	}
	ctx.JSON(tracks)
	return nil
}

func (Db *DB) GetTrack(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.SendString("Wrong id")
		return err
	}
	var track db.Track
	err = db.SelectTrack(Db.Db, &track, id)
	if err != nil {
		ctx.SendString("Cannot find track with given id")
		return err
	}
	ctx.JSONP(track)
	return nil
}

func (Db *DB) AddTrack(ctx *fiber.Ctx) error {
	var track db.Track
	if err := ctx.BodyParser(&track); err != nil {
		ctx.SendString("Wrong json struct")
		return err
	}
	if err := db.CreateTrack(Db.Db, &track); err != nil {
		ctx.SendString("Failed to add to Db")
		return err
	}
	ctx.SendString("Successfully added")
	ctx.JSON(track)
	return nil
}
func (Db *DB) UpdateTrack(ctx *fiber.Ctx) error {
	return nil
}

func (Db *DB) DeleteTrack(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		ctx.SendString("Wrong id")
		return err
	}
	var track db.Track
	err = db.SelectTrack(Db.Db, &track, id)
	if track.Artist == "" || err != nil {
		ctx.SendString("Cannot find track with given id")
		return err
	}
	err = db.DeleteTracks(Db.Db, &track, id)
	if err != nil {
		ctx.SendString("Cannot delete track with given id")
		return err
	}
	ctx.SendString("Successfully deleted")
	return nil
}
