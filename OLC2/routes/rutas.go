package routes

import (
	"OLC2/controllers"
	"github.com/gorilla/mux"
)

func RutasCompiladores(router *mux.Router) {
	router.HandleFunc("/", controllers.Inicio()).Methods("GET")
	router.HandleFunc("/prueba", controllers.Data()).Methods("POST")
}
