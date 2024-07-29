package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rnsasg/clean_architecture/internal/models"
	"github.com/rnsasg/clean_architecture/internal/repositories"
)

type TagRepo *repositories.TagRepository
type EntityRepo *repositories.EntityRepository
type TagEntityRepo *repositories.TagEntityRepository

type JSONLoaderService struct {
	TagRepo       *repositories.TagRepository
	EntityRepo    *repositories.EntityRepository
	TagEntityRepo *repositories.TagEntityRepository
}

func (s *JSONLoaderService) LoadTagsFromJSON(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var tags []models.Tag
	if err := json.NewDecoder(file).Decode(&tags); err != nil {
		return fmt.Errorf("failed to decode JSON file: %w", err)
	}

	for _, tag := range tags {
		if err := s.TagRepo.Create(&tag); err != nil {
			return fmt.Errorf("failed to save tag to database: %w", err)
		}
	}

	return nil
}

func (s *JSONLoaderService) LoadEntitiesFromJSON(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var entities []models.Entity
	if err := json.NewDecoder(file).Decode(&entities); err != nil {
		return fmt.Errorf("failed to decode JSON file: %w", err)
	}

	for _, entity := range entities {
		if err := s.EntityRepo.Create(&entity); err != nil {
			return fmt.Errorf("failed to save entity to database: %w", err)
		}
	}

	return nil
}

func (s *JSONLoaderService) LoadTagEntitiesFromJSON(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var tagEntities []models.TagEntity
	if err := json.NewDecoder(file).Decode(&tagEntities); err != nil {
		return fmt.Errorf("failed to decode JSON file: %w", err)
	}

	for _, tagEntity := range tagEntities {
		if err := s.TagEntityRepo.Create(&tagEntity); err != nil {
			return fmt.Errorf("failed to save tag_entity to database: %w", err)
		}
	}

	return nil
}

func (s *JSONLoaderService) LoadFromJSON() error {

	if err := s.LoadTagsFromJSON(TagJSONFile); err != nil {
		log.Fatalf("Error loading tags from JSON: %v", err)
		return err
	}
	if err := s.LoadEntitiesFromJSON(EntityJSONFile); err != nil {
		log.Fatalf("Error loading entities from JSON: %v", err)
		return err
	}
	if err := s.LoadTagEntitiesFromJSON(TagEntitiesJSONFile); err != nil {
		log.Fatalf("Error loading tag_entities from JSON: %v", err)
		return err
	}
	log.Println("Data loaded from JSON files.")
	return nil
}
