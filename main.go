package main

import (
	"todo-backend/database"
	"todo-backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET, POST, PUT, DELETE",
		},
	))
	app.Get("todos", handlers.GetTodos)
	app.Post("todos", handlers.CreateTodo)
	app.Delete("todos/:id", handlers.DeleteTodo)
	app.Put("todos/:id", handlers.UpdateTodo)

	app.Listen(":4000")
}
