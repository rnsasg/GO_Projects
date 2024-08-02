package repositories

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/rnsasg/Tagging/models"
)

type EntityRepository interface {
	GetAll() ([]models.Entity, error)
	GetByID(id gocql.UUID) (*models.Entity, error)
	Create(entiry *models.Entity) error
	Update(entiry *models.Entity) error
	Delete(id gocql.UUID) error
}

type entityRepository struct {
	session *gocql.Session
}

func NewEntityRepository(session *gocql.Session) EntityRepository {
	return &entityRepository{session}
}

func (r *entityRepository) GetAll() ([]models.Entity, error) {
	fmt.Println("GetAll")
	var entities []models.Entity
	iter := r.session.Query("SELECT id,name,type,metadata FROM entities").Iter()
	var entity models.Entity

	for iter.Scan(&entity.ID, &entity.Name, &entity.Type, &entity.Metadata) {
		entities = append(entities, entity)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *entityRepository) GetByID(id gocql.UUID) (*models.Entity, error) {
	fmt.Println("GetByID")
	var entity models.Entity
	if err := r.session.Query("SELECT id,name,type, matadata FROM entities WHERE id = ? ", id).Scan(&entity.ID, &entity.Name, &entity.Type, &entity.Metadata); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &entity, nil
}

func (r *entityRepository) Create(entity *models.Entity) error {
	fmt.Println("Create")
	return r.session.Query("INSERT INTO entities (id,name,type,metadata) VALUES (?,?,?,?)", entity.ID, entity.Name, entity.Type, entity.Metadata).Exec()
}

func (r *entityRepository) Update(entity *models.Entity) error {
	fmt.Println("Update")
	return r.session.Query("UPDATE entities SET name=?,type=?,metadata=? WHERE id = ? ", entity.Name, entity.Type, entity.Metadata, entity.ID).Exec()
}

func (r *entityRepository) Delete(id gocql.UUID) error {
	fmt.Println("Delete")
	return r.session.Query("DELETE FROM entities where id = ? ", id).Exec()
}
