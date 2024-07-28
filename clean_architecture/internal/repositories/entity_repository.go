package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/rnsasg/clean_architecture/internal/models"
	"os"
)

type EntityRepository struct {
	session *gocql.Session
}

func NewEntityRepository(session *gocql.Session) *EntityRepository {
	return &EntityRepository{
		session: session,
	}
}

func (r *EntityRepository) Create(entity *models.Entity) error {
	return r.session.Query(`INSERT INTO entities (id, name, type, metadata) VALUES (?, ?, ?, ?)`,
		entity.ID, entity.Name, entity.Type, entity.Metadata).Exec()
}

func (r *EntityRepository) GetAll() ([]models.Entity, error) {
	var entities []models.Entity
	iter := r.session.Query(`SELECT id, name, type, metadata FROM entities`).Iter()
	var entity models.Entity
	for iter.Scan(&entity.ID, &entity.Name, &entity.Type, &entity.Metadata) {
		entities = append(entities, entity)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return entities, nil
}

// SaveEntitiesToJSON saves all entities to a JSON file
func (r *EntityRepository) SaveEntitiesToJSON(filename string) error {
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
