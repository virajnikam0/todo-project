package main

import (
	"net/http"

	"github.com/virajnikam0/todo-project/web-todo-project/config"
	"github.com/virajnikam0/todo-project/youtubemp3/internal/api"
)

func main() {

	cfg := config.GetConfig()

	server := http.Server{
		Addr: cfg.Port,
	}

	handler := http.NewServeMux()



	// get the youtube link and verify it
	handler.HandleFunc("/getlink",api.GetYoutubeLink)







	server.Handler = handler

	server.ListenAndServe()


}