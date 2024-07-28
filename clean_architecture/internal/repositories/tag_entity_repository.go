package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/rnsasg/clean_architecture/internal/models"
	"os"
)

type TagEntityRepository struct {
	session *gocql.Session
}

func NewTagEntityRepository(session *gocql.Session) *TagEntityRepository {
	return &TagEntityRepository{
		session: session,
	}
}

func (r *TagEntityRepository) Create(tagEntity *models.TagEntity) error {
	return r.session.Query(`INSERT INTO tag_entities (tag_id, entity_id) VALUES (?, ?)`,
		tagEntity.TagID, tagEntity.EntityID).Exec()
}

func (r *TagEntityRepository) GetAll() ([]models.TagEntity, error) {
	var tagEntities []models.TagEntity
	iter := r.session.Query(`SELECT tag_id, entity_id FROM tag_entities`).Iter()
	var tagEntity models.TagEntity
	for iter.Scan(&tagEntity.TagID, &tagEntity.EntityID) {
		tagEntities = append(tagEntities, tagEntity)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return tagEntities, nil
}

// SaveTagEntitiesToJSON saves all tag-entity relationships to a JSON file
func (r *TagEntityRepository) SaveTagEntitiesToJSON(filename string) error {
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
