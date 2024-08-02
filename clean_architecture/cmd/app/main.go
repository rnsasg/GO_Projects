package main

import (
	"flag"
	"fmt"
	"github.com/rnsasg/clean_architecture/internal/db"
	"github.com/rnsasg/clean_architecture/internal/repositories"
	"github.com/rnsasg/clean_architecture/internal/services"
	"log"
)

const (
	datasize int = 10
)

func main() {
	// Command-line flags to choose between generating or loading data
	generateData := flag.Bool("generate", false, "Generate fresh data")
	loadFromJSON := flag.Bool("load_json", false, "Load data from JSON files")
	flag.Parse()

	// Initialize the Cassandra connection
	cassandra, err := db.NewCassandra()
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}
	defer cassandra.Close()

	// Initialize repositories
	tagRepo := repositories.NewTagRepository(*cassandra)
	entityRepo := repositories.NewEntityRepository(*cassandra)
	tagEntityRepo := repositories.NewTagEntityRepository(*cassandra)

	// Truncate DB Tables
	fmt.Println(cassandra.TruncateTable([]string{"entities", "tags", "tag_entities"}))

	if *generateData {
		dataGenerator := &services.DataGeneratorService{
			TagRepo:       tagRepo,
			EntityRepo:    entityRepo,
			TagEntityRepo: tagEntityRepo,
		}
		dataGenerator.GenerateData(*cassandra, datasize, true)

	} else if *loadFromJSON {
		// Load data from JSON files
		jsonLoader := &services.JSONLoaderService{
			TagRepo:       tagRepo,
			EntityRepo:    entityRepo,
			TagEntityRepo: tagEntityRepo,
		}
		jsonLoader.LoadFromJSON()
	} else {
		log.Println("No operation specified. Use -generate, -load_csv, or -load_json.")
	}
}
