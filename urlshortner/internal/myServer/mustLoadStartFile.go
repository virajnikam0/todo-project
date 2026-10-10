package myServer

import (
	"fmt"
	"net/http"

	"github.com/virajnikam0/todo-project/urlshortner/config"
	urlservice "github.com/virajnikam0/todo-project/urlshortner/internal/API/urlService"
)

// the function is responsible for sercver creation
func MustLoad() {

	// configuration
	cfg := config.GetConfig()

	// create the server
	server := http.Server{
		Addr: cfg.PORT,
	}

	// handler API
	handler := http.NewServeMux()
	/*
		-----------------------------------------------------
		-----------------------------------------------------
		-----------------------------------------------------
		-----------------------------------------------------
		-----------------------------------------------------
		-----------------------------------------------------
	*/

	// caling all the API here
	handler.HandleFunc("/hi", urlservice.GetLink)

	// pass the handler to server handler
	server.Handler = handler

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error during the server Listening ", err)
	}

}
