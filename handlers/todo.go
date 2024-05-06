package handlers

import (
	"go-htmx/types"
	"go-htmx/views/todo"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HandleGetTodo(c *fiber.Ctx, store []types.Task) error {
	return Render(c, c.Response().BodyWriter(), todo.Index(store))
}

func HandleGetTodolist(c *fiber.Ctx, store []types.Task) error {
	<-time.After(time.Second * 2) //To show the preloader :D
	return Render(c, c.Response().BodyWriter(), todo.List(store))
}

func HandlePostTodo(c *fiber.Ctx, store []types.Task) error {
	task := c.FormValue("task", "empty")
	newTask := types.Task{task, false}
	store = append(store, newTask)
	return Render(c, c.Response().BodyWriter(), todo.Task(newTask))
}
