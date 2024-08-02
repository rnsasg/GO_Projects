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

func (s *DataGeneratorService) GenerateData(cassandra db.Cassandra, size int, json bool) error {
	s.TagRepo.GenerateTags(TagJSONFile, size, json)
	s.EntityRepo.GenerateEntities(EntityJSONFile, size, json)
	s.TagEntityRepo.GenerateTagEntities(TagEntitiesJSONFile, size, json, s.TagRepo.GetAllTagIDs(), s.EntityRepo.GetAllEntityIDs())
	return nil
}
