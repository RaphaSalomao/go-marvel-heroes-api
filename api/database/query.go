package database

const (
	HeroesDDL = `
	CREATE TABLE IF NOT EXISTS heroes (
		id UUID PRIMARY KEY DEFAULT public.uuid_generate_v4(),
		marvel_id integer,
		power_strenght smallint,
		name VARCHAR(255),
		created_at timestamp DEFAULT now(),
		updated_at timestamp
	);`
	HeroesDescriptionDDL = `
	ALTER TABLE heroes ADD COLUMN IF NOT EXISTS description TEXT;`
)
