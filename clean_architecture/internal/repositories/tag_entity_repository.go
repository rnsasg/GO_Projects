package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/rnsasg/clean_architecture/internal/db"
	"github.com/rnsasg/clean_architecture/internal/models"
	"log"
	"math/rand"
	"os"
)

type TagEntityRepository struct {
	cassandra db.Cassandra
}

func NewTagEntityRepository(cassandra db.Cassandra) *TagEntityRepository {
	return &TagEntityRepository{
		cassandra: cassandra,
	}
}

func (r *TagEntityRepository) Create(tagEntity *models.TagEntity) error {
	return r.cassandra.ExecuteQuery(`INSERT INTO tag_entities (tag_id, entity_id) VALUES (?, ?)`,
		tagEntity.TagID, tagEntity.EntityID)
}

func (r *TagEntityRepository) GetAll() ([]models.TagEntity, error) {
	var tagEntities []models.TagEntity
	var tagEntity models.TagEntity

	iter := r.cassandra.Iterator(`SELECT tag_id, entity_id FROM tag_entities`)
	for iter.Scan(&tagEntity.TagID, &tagEntity.EntityID) {
		tagEntities = append(tagEntities, tagEntity)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return tagEntities, nil
}

// SaveTagEntitiesToJSONFromDB saves all tag-entity relationships to a JSON file
func (r *TagEntityRepository) SaveTagEntitiesToJSONFromDB(filename string) error {
	tagEntities, err := r.GetAll()
	if err != nil {
		return fmt.Errorf("failed to fetch tag-entity relationships: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tagEntities); err != nil {
		return fmt.Errorf("failed to encode tag-entity relationships to JSON: %v", err)
	}

	return nil
}

<<<<<<< HEAD
func (r *TagEntityRepository) GenerateTagEntities(filename string, size int, jsonflag bool, tagIDs, entityIDs []gocql.UUID) {

	var err error
	var jsonFile *os.File
	if jsonflag == true {
		jsonFile, err = os.Create(filename)
		if err != nil {
			log.Fatal("Could not create JSON file:", err)
		}
		defer jsonFile.Close()
	}
=======
func (r *TagEntityRepository) GenerateTagEntities(filename string, size int, tagIDs, entityIDs []gocql.UUID) {

	jsonFile, err := os.Create(filename)
	if err != nil {
		log.Fatal("Could not create JSON file:", err)
	}
	defer jsonFile.Close()
>>>>>>> 4128a271e7e9468cd9ffe65590ac94cbe867a96d

	var tagEntities []models.TagEntity

	// Generate and insert records into 'tag_entities' table
	for i := 0; i < size; i++ {
		tagID := tagIDs[rand.Intn(len(tagIDs))]
		entityID := entityIDs[rand.Intn(len(entityIDs))]

		// Insert into Cassandra
		if err := r.cassandra.ExecuteQuery(`INSERT INTO tag_entities (tag_id, entity_id) VALUES (?, ?)`,
			tagID, entityID); err != nil {
			log.Fatal(err)
		}

		// Append to JSON array
		tagEntities = append(tagEntities, models.TagEntity{
			TagID:    tagID,
			EntityID: entityID,
		})

		if i%1000 == 0 {
			fmt.Printf("Inserted and saved %d records\n", i+1)
		}
	}

<<<<<<< HEAD
	if jsonflag == true {
		// Write JSON file
		jsonData, err := json.Marshal(tagEntities)
		if err != nil {
			log.Fatal("Error marshalling data to JSON:", err)
		}
		jsonFile.Write(jsonData)
	}
=======
	// Write JSON file
	jsonData, err := json.Marshal(tagEntities)
	if err != nil {
		log.Fatal("Error marshalling data to JSON:", err)
	}
	jsonFile.Write(jsonData)
>>>>>>> 4128a271e7e9468cd9ffe65590ac94cbe867a96d

	fmt.Println("TagEntities data generation completed and saved to files.")
}
