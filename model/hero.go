package model

import (
	"time"

	"github.com/google/uuid"
)

const HeroesDDL = `
	CREATE TABLE IF NOT EXISTS heroes (
		id UUID PRIMARY KEY DEFAULT public.uuid_generate_v4(),
		marvel_id integer,
		power_strenght smallint,
		name VARCHAR(255),
		created_at timestamp DEFAULT now(),
		updated_at timestamp
	);`

type Hero struct {
	ID            uuid.UUID `json:"id,omitempty"`
	MarvelID      int       `json:"marvel_id,omitempty"`
	PowerStrenght int       `json:"power_strenght,omitempty"`
	Name          string    `json:"name,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

func (Hero) TableName() string {
	return "heroes"
}
