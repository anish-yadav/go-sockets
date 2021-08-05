package router

import (
	"log"

	"github.com/gorilla/mux"
)

func StartServer(logger log.Logger) *mux.Router {
	r := mux.NewRouter()
	addRoutes(r)
	return r
}
