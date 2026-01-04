package seeder

import (
	"github.com/sanskarchoudhry/pokedex-api-go/internal/database"
	"github.com/sanskarchoudhry/pokedex-api-go/internal/pokeapi"
)

// Seeder holds the dependencies needed to populate the DB
type Seeder struct {
	api *pokeapi.Client
	db  *database.Queries
}

func New(api *pokeapi.Client, db *database.Queries) *Seeder {
	return &Seeder{
		api: api,
		db:  db,
	}
}
