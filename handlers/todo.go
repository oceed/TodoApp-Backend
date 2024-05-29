package handlers

import (
	"todo-backend/database"
	"todo-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, title, status From todos")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if todo.Status == "" {
		todo.Status = "pending"
	}

	err := database.DB.QueryRow("INSERT INTO todos (title, status) VALUES ($1, $2) Returning id", todo.Title, todo.Status).Scan(&todo.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var todo models.Response
	todo.Message = "Delete Successfully"
	todo.Status = "Successfully"
	return c.Status(200).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := database.DB.Exec("UPDATE todos SET title = $1, status = $2 WHERE id = $3", todo.Title, todo.Status, id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(todo)
}
