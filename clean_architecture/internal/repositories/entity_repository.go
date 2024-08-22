package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/rnsasg/clean_architecture/internal/db"
	"github.com/rnsasg/clean_architecture/internal/models"
	"github.com/rnsasg/clean_architecture/utils"
	"log"
	"math/rand"
	"os"
	"time"
)

type EntityRepository struct {
	cassandra db.Cassandra
}

func NewEntityRepository(cassandra db.Cassandra) *EntityRepository {
	return &EntityRepository{
		cassandra: cassandra,
	}
}

func (r *EntityRepository) Create(entity *models.Entity) error {
	return r.cassandra.ExecuteQuery(`INSERT INTO entities (id, name, type, metadata) VALUES (?, ?, ?, ?)`,
		entity.ID, entity.Name, entity.Type, entity.Metadata)
}

func (r *EntityRepository) GetAll() ([]models.Entity, error) {
	var entities []models.Entity
	var entity models.Entity

	iter := r.cassandra.Iterator(`SELECT id, name, type, metadata FROM entities`)

	for iter.Scan(&entity.ID, &entity.Name, &entity.Type, &entity.Metadata) {
		entities = append(entities, entity)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return entities, nil
}

// SaveEntitiesToJSONFROMDB saves all entities to a JSON file
func (r *EntityRepository) SaveEntitiesToJSONFROMDB(filename string) error {
	entities, err := r.GetAll()
	if err != nil {
		return fmt.Errorf("failed to fetch entities: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(entities); err != nil {
		return fmt.Errorf("failed to encode entities to JSON: %v", err)
	}

	return nil
}

func (r *EntityRepository) GetAllEntityIDs() []gocql.UUID {
	var entityIDs []gocql.UUID
	iter := r.cassandra.Iterator(`SELECT id FROM entities`)
	var id gocql.UUID
	for iter.Scan(&id) {
		entityIDs = append(entityIDs, id)
	}
	if err := iter.Close(); err != nil {
		log.Fatalf("Failed to fetch entity IDs: %v", err)
	}
	return entityIDs
}

func (r *EntityRepository) GenerateEntities(filename string, size int, jsonflag bool) {
	rand.Seed(time.Now().UnixNano())

	var err error
	var jsonFile *os.File
	if jsonflag == true {
		jsonFile, err = os.Create(filename)
		if err != nil {
			log.Fatal("Could not create JSON file:", err)
		}
		defer jsonFile.Close()
	}

	var entities []models.Entity

	// Generate and insert records into 'entities' table
	for i := 0; i < size; i++ {
		entityID := gocql.TimeUUID()
		name := utils.RandomString(15)
		entityType := utils.RandomEntityType()
		metadata := utils.RandomString(10)

		// Insert into Cassandra
		if err := r.cassandra.ExecuteQuery(`INSERT INTO entities (id, name, type, metadata) VALUES (?, ?, ?, ?)`,
			entityID, name, entityType, metadata); err != nil {
			log.Fatal(err)
		}

		if jsonflag == true {
			// Append to JSON array
			entities = append(entities, models.Entity{
				ID:       entityID,
				Name:     name,
				Type:     entityType,
				Metadata: metadata,
			})
		}

		if i%1000 == 0 {
			fmt.Printf("Inserted and saved %d records\n", i+1)
		}
	}

	// Write JSON file
	if jsonflag == true {
		jsonData, err := json.Marshal(entities)
		if err != nil {
			log.Fatal("Error marshalling data to JSON:", err)
		}
		jsonFile.Write(jsonData)
	}

	fmt.Println("Data generation completed and saved to files.")
}
