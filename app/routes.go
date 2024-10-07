package app

import "github.com/fakh1m/LavanyaStore/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/Lavanya", controllers.Home).Methods("GET")
}
