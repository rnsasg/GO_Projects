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
)

type TagRepository struct {
	cassandra db.Cassandra
}

func NewTagRepository(cassandra db.Cassandra) *TagRepository {
	return &TagRepository{
		cassandra: cassandra,
	}
}

func (r *TagRepository) Create(tag *models.Tag) error {
	return r.cassandra.ExecuteQuery(`INSERT INTO tags (id, owner_id, name, description, properties) VALUES (?, ?, ?, ?, ?)`,
		tag.ID, tag.OwnerID, tag.Name, tag.Description, tag.Properties)
}

func (r *TagRepository) GetAll() ([]models.Tag, error) {
	var tags []models.Tag
	iter := r.cassandra.Iterator(`SELECT id, owner_id, name, description, properties FROM tags`)
	var tag models.Tag
	for iter.Scan(&tag.ID, &tag.OwnerID, &tag.Name, &tag.Description, &tag.Properties) {
		tags = append(tags, tag)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return tags, nil
}

// SaveTagsToJSONFromDB saves all tags to a JSON file
func (r *TagRepository) SaveTagsToJSONFromDB(filename string) error {
	tags, err := r.GetAll()
	if err != nil {
		return fmt.Errorf("failed to fetch tags: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tags); err != nil {
		return fmt.Errorf("failed to encode tags to JSON: %v", err)
	}

	return nil
}

func (r *TagRepository) GetAllTagIDs() []gocql.UUID {
	var tagIDs []gocql.UUID
	iter := r.cassandra.Iterator(`SELECT id FROM tags`)
	var id gocql.UUID
	for iter.Scan(&id) {
		tagIDs = append(tagIDs, id)
	}
	if err := iter.Close(); err != nil {
		log.Fatalf("Failed to fetch tag IDs: %v", err)
	}
	return tagIDs
}

func (r *TagRepository) GenerateTags(filename string, size int, jsonflag bool) {

	var err error
	var jsonFile *os.File
	if jsonflag == true {
		jsonFile, err = os.Create(filename)
		if err != nil {
			log.Fatal("Could not create JSON file:", err)
		}
		defer jsonFile.Close()
	}

	var tags []models.Tag

	// Generate and insert records into 'tags' table
	for i := 0; i < size; i++ {
		tagID := gocql.TimeUUID()
		ownerID := rand.Intn(1000) + 1
		name := utils.RandomString(15)
		description := utils.RandomString(50)
		properties := utils.RandomString(100)

		// Insert into Cassandra
		if err := r.cassandra.ExecuteQuery(`INSERT INTO tags (id, owner_id, name, description, properties) VALUES (?, ?, ?, ?, ?)`,
			tagID, ownerID, name, description, properties); err != nil {
			log.Fatal(err)
		}

		if jsonflag == true {
			// Append to JSON array
			tags = append(tags, models.Tag{
				ID:          tagID,
				OwnerID:     ownerID,
				Name:        name,
				Description: description,
				Properties:  properties,
			})
		}

		if i%1000 == 0 {
			fmt.Printf("Inserted and saved %d records\n", i+1)
		}
	}

	if jsonflag == true {
		// Write JSON file
		jsonData, err := json.Marshal(tags)
		if err != nil {
			log.Fatal("Error marshalling data to JSON:", err)
		}
		jsonFile.Write(jsonData)
	}

	fmt.Println("Tags data generation completed and saved to files.")
}
