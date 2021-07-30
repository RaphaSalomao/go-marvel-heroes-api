package api

import (
	"br.com.github/raphasalomao/go-marvel-heroes-api/api/controller"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func InitController() {
	Router = mux.NewRouter()
	Router.HandleFunc("/api/v1/hero", controller.CreateHero).Methods("POST")
	Router.HandleFunc("/api/v1/hero/{id}", controller.FindHeroById).Methods("GET")
	Router.HandleFunc("/api/v1/hero", controller.FindAllHeroes).Methods("GET")
	Router.HandleFunc("/api/v1/hero/{id}", controller.DeleteHeroById).Methods("DELETE")
	Router.HandleFunc("/api/v1/hero", controller.UpdateHero).Methods("PATCH")
}
