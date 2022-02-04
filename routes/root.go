package routes

import "github.com/gofiber/fiber/v2"

func serverCheck(ctx *fiber.Ctx) error {
	return ctx.SendString("Server is Running at port 8080")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", serverCheck)
	app.Post("/create/task", createTask)
	app.Get("/list/task", listTasks)
	app.Delete("/delete/task/:id", deleteTasks)
	app.Put("/update/task/:id", updateTask)
	app.Get("/search/task", searchTasks)
}
