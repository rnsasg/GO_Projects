package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gocql/gocql"
	"github.com/rnsasg/clean_architecture/internal/models"
	"github.com/rnsasg/clean_architecture/internal/repositories"
)

const (
	TotalNoOfRecords = 10
)

type DataGeneratorService struct {
	TagRepo       *repositories.TagRepository
	EntityRepo    *repositories.EntityRepository
	TagEntityRepo *repositories.TagEntityRepository
}

func randomString(length int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, length)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}

func (s *DataGeneratorService) GenerateData() error {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()

	// Generate and insert 1 million records into 'tags' table
	for i := 0; i < TotalNoOfRecords; i++ {
		tagID, _ := gocql.ParseUUID(gocql.TimeUUID().String())
		ownerID := rand.Intn(1000) + 1
		name := randomString(15)
		description := randomString(50)
		properties := randomString(100)
		tag := &models.Tag{
			ID:          tagID,
			OwnerID:     ownerID,
			Name:        name,
			Description: description,
			Properties:  properties,
		}
		if err := s.TagRepo.Create(tag); err != nil {
			return err
		}
		if i%1000 == 0 {
			fmt.Printf("Inserted %d records into 'tags'\n", i+1)
		}
	}

	tagsDuration := time.Since(startTime)
	fmt.Printf("Time taken to generate and insert 1 million records into 'tags': %s\n", tagsDuration)

	startTime = time.Now()

	// Generate and insert 1 million records into 'entities' table
	for i := 0; i < TotalNoOfRecords; i++ {
		entityID, _ := gocql.ParseUUID(gocql.TimeUUID().String())
		name := randomString(15)
		entityType := randomString(10)
		metadata := randomString(100)
		entity := &models.Entity{
			ID:       entityID,
			Name:     name,
			Type:     entityType,
			Metadata: metadata,
		}
		if err := s.EntityRepo.Create(entity); err != nil {
			return err
		}
		if i%1000 == 0 {
			fmt.Printf("Inserted %d records into 'entities'\n", i+1)
		}
	}

	entitiesDuration := time.Since(startTime)
	fmt.Printf("Time taken to generate and insert 1 million records into 'entities': %s\n", entitiesDuration)

	startTime = time.Now()

	// Select all tag IDs and entity IDs to use in the tag_entities table
	tagIDs, err := s.TagRepo.GetAll()
	if err != nil {
		return err
	}
	entityIDs, err := s.EntityRepo.GetAll()
	if err != nil {
		return err
	}

	// Generate and insert records into 'tag_entities' table using IDs from 'tags' and 'entities'
	for i := 0; i < TotalNoOfRecords; i++ {
		tagID := tagIDs[rand.Intn(len(tagIDs))].ID
		entityID := entityIDs[rand.Intn(len(entityIDs))].ID
		tagEntity := &models.TagEntity{
			TagID:    tagID,
			EntityID: entityID,
		}
		if err := s.TagEntityRepo.Create(tagEntity); err != nil {
			return err
		}
		if i%1000 == 0 {
			fmt.Printf("Inserted %d records into 'tag_entities'\n", i+1)
		}
	}

	tagEntitiesDuration := time.Since(startTime)
	fmt.Printf("Time taken to generate and insert records into 'tag_entities': %s\n", tagEntitiesDuration)

	totalDuration := tagsDuration + entitiesDuration + tagEntitiesDuration
	fmt.Printf("Total time taken for data generation and insertion: %s\n", totalDuration)

	return nil
}
