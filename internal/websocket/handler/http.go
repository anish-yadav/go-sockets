package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type SocketHandler struct {
}

func NewSocketHanadler() *SocketHandler {
	return &SocketHandler{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(ws *websocket.Conn) {
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error ws %s \n", err.Error())
			return
		}

		log.Printf("Message %s\n", string(p))
		if err = ws.WriteMessage(messageType, p); err != nil {
			log.Printf("Error send %s \n", err.Error())
			return
		}
	}
}

func (s *SocketHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage")
}

func (s *SocketHandler) WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error %s\n", err.Error())
	}
	reader(ws)
	log.Println("Successfully connected")
}
