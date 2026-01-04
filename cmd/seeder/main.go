package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // The Postgres Driver (Blank import is intentional)
	"github.com/sanskarchoudhry/pokedex-api-go/internal/pokeapi"
)

func main() {
	// 1. Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Get DB URL
	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		log.Fatal("DB_SOURCE is not set in .env")
	}

	// 3. Connect to Database
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("Could not open connection to DB:", err)
	}
	defer db.Close() // Close connection when main function finishes

	// 4. Ping to verify connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not ping DB:", err)
	}

	fmt.Println("✅ Successfully connected to Pokedex Database!")

	pokeClient := pokeapi.NewClient("https://pokeapi.co/api/v2", 5*time.Second)

	gens, err := pokeClient.ListGenerations()
	if err != nil {
		log.Fatal("Error fetching generations:", err)
	}

	fmt.Printf("Fetched %d generations\n", gens.Count)
	for _, g := range gens.Results {
		fmt.Println(" - " + g.GenName)
	}
}
