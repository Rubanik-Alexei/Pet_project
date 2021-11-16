package main

import (
	"Pet_project/db"
	"Pet_project/handlers"
	"fmt"

	"github.com/gofiber/fiber"
)

func CreateHandlers(app *fiber.App) {
	app.Get("/track", handlers.GetAlltracks)
	app.Post("/add", handlers.AddTrack)
	app.Get("/track/:id", handlers.GetTrack)
	app.Delete("/track/:id", handlers.DeleteTrack)
	app.Post("/track/:id", handlers.UpdateTrack)
}

func main() {
	fiberApp := fiber.New()
	dbConn, err := db.ConnectToDB()
	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Connected to DB")
	fiberApp.Listen(9000)
	defer dbConn.Close()
}
