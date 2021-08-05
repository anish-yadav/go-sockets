package router

import (
	"github.com/anish-yadav/go-sockets/internal/websocket/handler"
	"github.com/gorilla/mux"
)

func addRoutes(r *mux.Router) {
	sh := handler.NewSocketHanadler()
	r.HandleFunc("/ws", sh.WsEndpoint)
	r.HandleFunc("/", sh.HomePage)
}
