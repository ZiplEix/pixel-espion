package routes

import (
	"github.com/ZiplEix/pixel-espion/controllers"
	"github.com/ZiplEix/pixel-espion/middlewares"
	"github.com/gofiber/fiber/v2"
)

func spyRoutes(app *fiber.App) {
	app.Get("/spy/pixel1", controllers.Pixel1)

	spyGroup := app.Group("/spy", middlewares.Protected)
	spyGroup.Post("/new", controllers.NewSpy)
	spyGroup.Get("/all", controllers.GetAllSpies)
	spyGroup.Get("/:id", controllers.GetSpy)
	spyGroup.Put("/:id", controllers.UpdateSpy)
	spyGroup.Delete("/:id", controllers.DeleteSpy)

	recordGroup := app.Group("/record", middlewares.Protected)
	recordGroup.Get("/all", controllers.GetAllRecords)
	recordGroup.Get("/spy/:id", controllers.GetSpyRecords)
	recordGroup.Delete("/:id", controllers.DeleteRecord)
}
