package main

import (
	"fmt"
	"github.com/anish-yadav/go-sockets/internal/websocket/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Websockets")
	logger := log.Default()
	r := router.StartServer(*logger)

	serv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	if err := serv.ListenAndServe(); err != nil {
		log.Printf("Error  %s \n", err.Error())
	}
}
