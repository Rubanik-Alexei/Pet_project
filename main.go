package main

import (
	"Pet_project/handlers"

	"github.com/gofiber/fiber/v2"
)

/*Commands for tests
curl -X POST -H "Content-Type: application/json" --data "{\"Artist\":\"Bones\",\"Name\":\"Dirt\",\"Genre\":\"Cloud\"}"  http://localhost:9000/add

*/

func CreateHandlers(handler *handlers.DB, app *fiber.App) {
	app.Get("/track", handler.GetAlltracks)
	app.Post("/add", handler.AddTrack)
	app.Get("/track/:id", handler.GetTrack)
	app.Delete("/track/:id", handler.DeleteTrack)
	app.Post("/track/:id", handler.UpdateTrack)
}

func main() {
	fiberApp := fiber.New()
	handler := handlers.NewDB()
	defer handler.Db.Close()
	CreateHandlers(handler, fiberApp)
	fiberApp.Listen("localhost:9000")
}
