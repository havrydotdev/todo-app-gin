package main

import (
	"log"

	"github.com/gavrylenkoIvan/todo-app-gin"
	"github.com/gavrylenkoIvan/todo-app-gin/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while initializing server, %s", err.Error())
	}
}
