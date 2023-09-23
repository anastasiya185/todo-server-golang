package main

import (
	"log"

	"github.com/anastasiya185/todo-server-golang"
	"github.com/anastasiya185/todo-server-golang/pkg/handler"
)

func main (){
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil{
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}