package services

import (
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/rnsasg/Tagging/models"
	"github.com/rnsasg/Tagging/repositories"
)

type EntityService interface {
	GetAllEntities() ([]models.Entity, error)
	GetEntityByID(id gocql.UUID) (*models.Entity, error)
	CreateEntity(entity *models.Entity) error
	UpdateEntity(entity *models.Entity) error
	DeleteEntity(id gocql.UUID) error
}

type entityService struct {
	entityRepository repositories.EntityRepository
}

func NewEntityService(entityRepo repositories.EntityRepository) EntityService {
	return &entityService{entityRepo}
}

func (s *entityService) GetAllEntities() ([]models.Entity, error) {
	return s.entityRepository.GetAll()
}

func (s *entityService) GetEntityByID(id gocql.UUID) (*models.Entity, error) {
	return s.entityRepository.GetByID(id)
}

func (s *entityService) CreateEntity(entity *models.Entity) error {
	newUUID := uuid.New()
	gocql.ParseUUID(newUUID.String())
	return s.entityRepository.Create(entity)
}

func (s *entityService) UpdateEntity(entity *models.Entity) error {
	return s.entityRepository.Update(entity)
}

func (s *entityService) DeleteEntity(id gocql.UUID) error {
	return s.entityRepository.Delete(id)
}
