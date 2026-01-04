package seeder

import (
	"context"
	"fmt"
	"log"

	"github.com/sanskarchoudhry/pokedex-api-go/internal/database"
)

func (s *Seeder) SeedGenerations() {
	fmt.Println("🌱 Seeding Generations...")

	// Use s.api (the struct field), not a new client
	gens, err := s.api.ListGenerations()
	if err != nil {
		log.Fatal("Error fetching generations:", err)
	}

	for _, g := range gens.Results {
		fmt.Println("Fetching details for generation: " + g.GenName)

		// Use s.api
		genDetails, err := s.api.GetGeneration(g.GenName)
		if err != nil {
			log.Printf("Error fetching details for %s: %v", g.GenName, err)
			continue
		}

		// Use s.db (the struct field)
		savedGen, err := s.db.CreateGeneration(context.Background(), database.CreateGenerationParams{
			ID:         int32(genDetails.ID),
			Name:       genDetails.Name,
			RegionName: genDetails.MainRegion.Name,
		})

		if err != nil {
			log.Printf("Error saving %s to DB: %v", genDetails.Name, err)
		} else {
			fmt.Printf("✅ Saved: %s (Region: %s)\n", savedGen.Name, savedGen.RegionName)
		}
	}
}
