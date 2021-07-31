package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"br.com.github/raphasalomao/go-marvel-heroes-api/api/client"
	"br.com.github/raphasalomao/go-marvel-heroes-api/api/database"
	"br.com.github/raphasalomao/go-marvel-heroes-api/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateHero(w http.ResponseWriter, r *http.Request) {
	var request model.CreateHeroRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Timestamp: time.Now(),
			Message:   err.Error(),
		})
		return
	}
	marvelHero, err := client.GetCharacters(request.Name)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	hero := model.Hero{
		ID:            uuid.New(),
		Name:          request.Name,
		PowerStrenght: request.PowerStrenght,
		MarvelID:      marvelHero.Data.Results[0].ID,
		Description:   marvelHero.Data.Results[0].Description,
	}
	database.DB.Create(&hero)
	w.Header().Set("Content-Location", fmt.Sprintf("/api/v1/hero/%s", hero.ID))
	w.WriteHeader(http.StatusCreated)
}

func UpdateHero(w http.ResponseWriter, r *http.Request) {
	var hero model.Hero
	err := json.NewDecoder(r.Body).Decode(&hero)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hero.UpdatedAt = time.Now()
	database.DB.Save(hero)
	w.WriteHeader(http.StatusNoContent)
}

func FindHeroById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var hero model.Hero
	result := database.DB.First(&hero, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(hero)
}

func FindAllHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	var heroes []model.Hero
	database.DB.Find(&heroes)
	json.NewEncoder(w).Encode(heroes)
}

func DeleteHeroById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	result := database.DB.Where("id = ?", id).Delete(&model.Hero{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
