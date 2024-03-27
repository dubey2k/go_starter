package routes

import (
	"github.com/dubey2k/go_starter/internal/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/getData", controllers.GetData).Methods("GET")
}
