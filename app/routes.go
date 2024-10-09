package app

import (
	"github.com/fakh1m/LavanyaStore/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/Lavanya", controllers.Home).Methods("GET")
}
