package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sanskarchoudhry/pokedex-api-go/internal/database"
	"github.com/sanskarchoudhry/pokedex-api-go/internal/pokeapi"
	"github.com/sanskarchoudhry/pokedex-api-go/internal/seeder"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbSource := os.Getenv("DB_SOURCE")
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("Could not open DB:", err)
	}
	defer db.Close()

	queries := database.New(db)
	pokeClient := pokeapi.NewClient("https://pokeapi.co/api/v2", 5*time.Second)

	sd := seeder.New(pokeClient, queries)

	sd.SeedGenerations()
	sd.SeedTypes()

	fmt.Println("✅ Seeding complete!")
}
