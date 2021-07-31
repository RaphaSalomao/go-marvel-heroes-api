package model

import (
	"time"

	"github.com/google/uuid"
)

type Hero struct {
	ID            uuid.UUID `json:"id,omitempty"`
	MarvelID      int       `json:"marvel_id,omitempty"`
	PowerStrenght int       `json:"power_strenght,omitempty"`
	Name          string    `json:"name,omitempty"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type CreateHeroRequest struct {
	PowerStrenght int    `json:"power_strenght,omitempty"`
	Name          string `json:"name,omitempty"`
}

type MarvelHeroResponse struct {
	Data struct {
		Results []struct {
			ID          int    `json:"id"`
			Description string `json:"description"`
		} `json:"results"`
	} `json:"data"`
}

func (Hero) TableName() string {
	return "heroes"
}
