package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // The Postgres Driver (Blank import is intentional)
	"github.com/sanskarchoudhry/pokedex-api-go/internal/database"
	"github.com/sanskarchoudhry/pokedex-api-go/internal/pokeapi"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		log.Fatal("DB_SOURCE is not set in .env")
	}

	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("Could not open connection to DB:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not ping DB:", err)
	}

	fmt.Println("✅ Successfully connected to Pokedex Database!")

	queries := database.New(db)

	pokeClient := pokeapi.NewClient("https://pokeapi.co/api/v2", 5*time.Second)

	gens, err := pokeClient.ListGenerations()
	if err != nil {
		log.Fatal("Error fetching generations:", err)
	}

	fmt.Printf("Fetched %d generations\n", gens.Count)
	for _, g := range gens.Results {
		fmt.Println("Fething details for generation" + g.GenName)

		genDetails, err := pokeClient.GetGeneration(g.GenName)
		if err != nil {
			log.Printf("Error fetching details for %s: %v", g.GenName, err)
			continue // Skip this one, don't crash the whole app
		}

		// Insert into DB
		savedGen, err := queries.CreateGeneration(context.Background(), database.CreateGenerationParams{
			ID:         int32(genDetails.ID), // explicit cast is good!
			Name:       genDetails.Name,
			RegionName: genDetails.MainRegion.Name,
		})

		if err != nil {
			log.Printf("Error saving %s to DB: %v", genDetails.Name, err)
		} else {
			fmt.Printf("✅ Saved: %s (Region: %s)\n", savedGen.Name, savedGen.RegionName)
		}
	}

	types, err := pokeClient.ListTypes()
	if err != nil {
		log.Fatal("Error fetching types", err)
	}

	for _, t := range types.Results {
		fmt.Println("Fetching details for type" + t.TypeName)

		typedetails, err := pokeClient.GetType(t.TypeName)
		if err != nil {
			log.Printf("Error fetching details for %s: %v", t.TypeName, err)
			continue
		}

		savedType, err := queries.CreateType(context.Background(), database.CreateTypeParams{
			ID:   int32(typedetails.ID),
			Name: typedetails.Name,
		})

		if err != nil {
			log.Printf("Error saving %s to DB: %v", typedetails.Name, err)
		} else {
			fmt.Printf("✅ Saved: %s\n", savedType.Name)
		}
	}
}
