package seeder

import (
	"context"
	"fmt"
	"log"

	"github.com/sanskarchoudhry/pokedex-api-go/internal/database"
)

func (s *Seeder) SeedTypes() {
	fmt.Println("🌱 Seeding Types...")

	typesList, err := s.api.ListTypes()
	if err != nil {
		log.Fatal("Error fetching types:", err)
	}

	for _, t := range typesList.Results {
		fmt.Println("Fetching details for type: " + t.TypeName)

		typeDetails, err := s.api.GetType(t.TypeName)
		if err != nil {
			log.Printf("Error fetching details for %s: %v", t.TypeName, err)
			continue
		}

		savedType, err := s.db.CreateType(context.Background(), database.CreateTypeParams{
			ID:   int32(typeDetails.ID),
			Name: typeDetails.Name,
		})

		if err != nil {
			log.Printf("Error saving %s to DB: %v", typeDetails.Name, err)
		} else {
			fmt.Printf("✅ Saved: %s\n", savedType.Name)
		}
	}
}
