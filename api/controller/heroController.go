package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"br.com.github/raphasalomao/go-marvel-heroes-api/api/database"
	"br.com.github/raphasalomao/go-marvel-heroes-api/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateHero(w http.ResponseWriter, r *http.Request) {
	var hero model.Hero
	err := json.NewDecoder(r.Body).Decode(&hero)
	if err != nil {
		log.Panic("Failed to parse request", err)
	}
	hero.ID = uuid.New()
	database.DB.Create(&hero)
}

func UpdateHero(w http.ResponseWriter, r *http.Request) {
	var hero model.Hero
	err := json.NewDecoder(r.Body).Decode(&hero)
	if err != nil {
		log.Panic("Failed to parse request", err)
	}
	hero.UpdatedAt = time.Now()
	database.DB.Save(hero)
	w.WriteHeader(http.StatusNoContent)
}

func FindHeroById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := uuid.Parse(params["id"])
	if err != nil {
		log.Panic("failed to parse request", err)
	}
	fmt.Println(id)
	var hero model.Hero
	database.DB.First(&hero, id)
	if hero.ID != uuid.Nil {
		json.NewEncoder(w).Encode(hero)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func FindAllHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	var heroes []model.Hero
	database.DB.Find(&heroes)
	json.NewEncoder(w).Encode(heroes)
}

func DeleteHeroById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	database.DB.Where("id = ?", id).Delete(&model.Hero{})
	w.WriteHeader(http.StatusNoContent)
}
