package main

import (
	"fmt"
	"github.com/ameykpatil/football-engine/engine"
	"github.com/ameykpatil/football-engine/utils/server"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		server.SendResponse(w, map[string]string{"message": "pong"}, http.StatusOK)
	})

	http.HandleFunc("/players/fetch", PlayerHandler)

	fmt.Println("Starting a server at localhost:4000")

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic(err)
	}
}

// PlayerHandler is a handler function for players fetch API
func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	engineInstance := engine.NewEngine()
	resp := engineInstance.Start()
	server.SendResponse(w, resp, http.StatusOK)
}
