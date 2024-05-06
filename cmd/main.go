package main

import (
	"go-htmx/handlers"
	"go-htmx/types"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func MakeFiberHandler(fn func(*fiber.Ctx, []types.Task) error, store []types.Task) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return fn(c, store)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	listenAddr := os.Getenv("LISTEN_ADDR")
	router := fiber.New(fiber.Config{
		Immutable: true,
	})

	store := []types.Task{
		types.Task{"test", true},
		types.Task{"test2", false},
	}

	router.Get("/", handlers.HandleGetIndex)
	router.Post("/login", handlers.HandlePostLogin)
	router.Get("/todo", MakeFiberHandler(handlers.HandleGetTodo, store))
	router.Get("/todolist", MakeFiberHandler(handlers.HandleGetTodolist, store))
	router.Post("/todo", MakeFiberHandler(handlers.HandlePostTodo, store))
	router.Use("/static", filesystem.New(filesystem.Config{
		Root: http.Dir("./static"),
	}))

	logrus.Info("server is running at port ", listenAddr)
	logrus.Error(router.Listen(listenAddr))
}
