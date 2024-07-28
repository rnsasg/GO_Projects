package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/rnsasg/clean_architecture/internal/models"
	"os"
)

type TagRepository struct {
	session *gocql.Session
}

func NewTagRepository(session *gocql.Session) *TagRepository {
	return &TagRepository{
		session: session,
	}
}

func (r *TagRepository) Create(tag *models.Tag) error {
	return r.session.Query(`INSERT INTO tags (id, owner_id, name, description, properties) VALUES (?, ?, ?, ?, ?)`,
		tag.ID, tag.OwnerID, tag.Name, tag.Description, tag.Properties).Exec()
}

func (r *TagRepository) GetAll() ([]models.Tag, error) {
	var tags []models.Tag
	iter := r.session.Query(`SELECT id, owner_id, name, description, properties FROM tags`).Iter()
	var tag models.Tag
	for iter.Scan(&tag.ID, &tag.OwnerID, &tag.Name, &tag.Description, &tag.Properties) {
		tags = append(tags, tag)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return tags, nil
}

// SaveTagsToJSON saves all tags to a JSON file
func (r *TagRepository) SaveTagsToJSON(filename string) error {
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
