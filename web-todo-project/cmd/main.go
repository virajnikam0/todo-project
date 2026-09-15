package main

import (
	"log"
	"net/http"

	"github.com/virajnikam0/todo-project/web-todo-project/config"
	myTask "github.com/virajnikam0/todo-project/web-todo-project/internal/api"
)

func main() {

	cfg := config.GetConfig()	
	// make the server 
	server := http.Server{
		Addr: cfg.Port,
	}
	// make the handler
	handler := http.NewServeMux()

	// create the API endpoint 
	/*
		1. add the todo
		2. delete the todo
		3. update the todo
		4. get all todos
		5. done and undone 

	*/
	
	handler.HandleFunc("/add",myTask.AddTask)
	handler.HandleFunc("/read",myTask.ReadTask)
	handler.HandleFunc("/delete/{id}",myTask.DeleteTask)
	handler.HandleFunc("/update/{id}",myTask.UpdateTask)




	server.Handler = handler
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server is not running: %v\n",err)
	}
}