package main

import (
	"flag"
	"github.com/rnsasg/clean_architecture/internal/db"
	"github.com/rnsasg/clean_architecture/internal/repositories"
	"github.com/rnsasg/clean_architecture/internal/services"
	"log"
)

func main() {
	// Command-line flags to choose between generating or loading data
	generateData := flag.Bool("generate", false, "Generate fresh data")
	loadFromCSV := flag.Bool("load_csv", false, "Load data from CSV files")
	loadFromJSON := flag.Bool("load_json", false, "Load data from JSON files")
	flag.Parse()

	// Initialize the Cassandra connection
	cassandra, err := db.NewCassandra()
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}
	defer cassandra.Close()

	// Initialize repositories
	tagRepo := repositories.NewTagRepository(cassandra.Session)
	entityRepo := repositories.NewEntityRepository(cassandra.Session)
	tagEntityRepo := repositories.NewTagEntityRepository(cassandra.Session)

	if *generateData {
		// Generate fresh data
		dataGenerator := &services.DataGeneratorService{
			TagRepo:       tagRepo,
			EntityRepo:    entityRepo,
			TagEntityRepo: tagEntityRepo,
		}
		if err := dataGenerator.GenerateData(); err != nil {
			log.Fatalf("Error generating data: %v", err)
		}
		log.Println("Data generation completed.")
	} else if *loadFromCSV {
		// Load data from CSV files
		csvLoader := &services.CSVLoaderService{
			TagRepo:       tagRepo,
			EntityRepo:    entityRepo,
			TagEntityRepo: tagEntityRepo,
		}
		if err := csvLoader.LoadTagsFromCSV("tags.csv"); err != nil {
			log.Fatalf("Error loading tags from CSV: %v", err)
		}
		if err := csvLoader.LoadEntitiesFromCSV("entities.csv"); err != nil {
			log.Fatalf("Error loading entities from CSV: %v", err)
		}
		if err := csvLoader.LoadTagEntitiesFromCSV("tag_entities.csv"); err != nil {
			log.Fatalf("Error loading tag_entities from CSV: %v", err)
		}
		log.Println("Data loaded from CSV files.")
	} else if *loadFromJSON {
		// Load data from JSON files
		jsonLoader := &services.JSONLoaderService{
			TagRepo:       tagRepo,
			EntityRepo:    entityRepo,
			TagEntityRepo: tagEntityRepo,
		}
		if err := jsonLoader.LoadTagsFromJSON("tags.json"); err != nil {
			log.Fatalf("Error loading tags from JSON: %v", err)
		}
		if err := jsonLoader.LoadEntitiesFromJSON("entities.json"); err != nil {
			log.Fatalf("Error loading entities from JSON: %v", err)
		}
		if err := jsonLoader.LoadTagEntitiesFromJSON("tag_entities.json"); err != nil {
			log.Fatalf("Error loading tag_entities from JSON: %v", err)
		}
		log.Println("Data loaded from JSON files.")
	} else {
		log.Println("No operation specified. Use -generate, -load_csv, or -load_json.")
	}
}
