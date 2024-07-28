package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gocql/gocql"
	"github.com/rnsasg/clean_architecture/internal/models"
	"github.com/rnsasg/clean_architecture/internal/repositories"
)

type CSVLoaderService struct {
	TagRepo       *repositories.TagRepository
	EntityRepo    *repositories.EntityRepository
	TagEntityRepo *repositories.TagEntityRepository
}

func (s *CSVLoaderService) LoadTagsFromCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	for _, record := range records {
		tagID, _ := gocql.ParseUUID(record[0])
		ownerID, _ := strconv.Atoi(record[1])
		tag := &models.Tag{
			ID:          tagID,
			OwnerID:     ownerID,
			Name:        record[2],
			Description: record[3],
			Properties:  record[4],
		}
		if err := s.TagRepo.Create(tag); err != nil {
			return fmt.Errorf("failed to save tag to database: %w", err)
		}
	}

	return nil
}

func (s *CSVLoaderService) LoadEntitiesFromCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	for _, record := range records {
		entityID, _ := gocql.ParseUUID(record[0])
		entity := &models.Entity{
			ID:       entityID,
			Name:     record[1],
			Type:     record[2],
			Metadata: record[3],
		}
		if err := s.EntityRepo.Create(entity); err != nil {
			return fmt.Errorf("failed to save entity to database: %w", err)
		}
	}

	return nil
}

func (s *CSVLoaderService) LoadTagEntitiesFromCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	for _, record := range records {
		tagID, _ := gocql.ParseUUID(record[0])
		entityID, _ := gocql.ParseUUID(record[1])
		tagEntity := &models.TagEntity{
			TagID:    tagID,
			EntityID: entityID,
		}
		if err := s.TagEntityRepo.Create(tagEntity); err != nil {
			return fmt.Errorf("failed to save tag_entity to database: %w", err)
		}
	}

	return nil
}
