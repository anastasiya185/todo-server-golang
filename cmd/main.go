package main

import (
	todo "_/Users/anastacia/Desktop/Go/Rest_server"
	"log"
)

func main (){
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil{
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}