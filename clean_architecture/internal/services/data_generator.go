package services

import (
	"github.com/rnsasg/clean_architecture/internal/db"
	"github.com/rnsasg/clean_architecture/internal/repositories"
)

type DataGeneratorService struct {
	TagRepo       *repositories.TagRepository
	EntityRepo    *repositories.EntityRepository
	TagEntityRepo *repositories.TagEntityRepository
}

func (s *DataGeneratorService) GenerateData(cassandra db.Cassandra, size int) error {
	s.TagRepo.GenerateTags(TagJSONFile, size)
	s.EntityRepo.GenerateEntities(EntityJSONFile, size)
	s.TagEntityRepo.GenerateTagEntities(TagEntitiesJSONFile, size, s.TagRepo.GetAllTagIDs(cassandra.Session), s.EntityRepo.GetAllEntityIDs(cassandra.Session))
	return nil
}
