package tests

import (
	"fmt"
	"github.com/rnsasg/clean_architecture/internal/db"
	"github.com/rnsasg/clean_architecture/internal/repositories"
	"github.com/rnsasg/clean_architecture/internal/services"
	"github.com/rnsasg/clean_architecture/utils"
	"log"
	"sync"
	"testing"
	"time"
)

func TestDBSerialOperationPerformance(t *testing.T) {

	// Initialize the Cassandra connection
	cassandra, err := db.NewCassandra()
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}

	defer func() {
		// Truncate DB Tables
		fmt.Println(cassandra.TruncateTable([]string{"entities", "tags", "tag_entities"}))
		// Close the connection
		cassandra.Close()
	}()

	// Initialize repositories
	tagRepo := repositories.NewTagRepository(*cassandra)
	entityRepo := repositories.NewEntityRepository(*cassandra)
	tagEntityRepo := repositories.NewTagEntityRepository(*cassandra)

	// Truncate DB Tables
	fmt.Println(cassandra.TruncateTable([]string{"entities", "tags", "tag_entities"}))

	dataGenerator := &services.DataGeneratorService{
		TagRepo:       tagRepo,
		EntityRepo:    entityRepo,
		TagEntityRepo: tagEntityRepo,
	}

	dataSizes := []int{10}
	for _, size := range dataSizes {
		t.Run(fmt.Sprintf("%d Records", size), func(t *testing.T) {
			start := time.Now()
			if err := dataGenerator.GenerateData(*cassandra, size, false); err != nil {
				t.Errorf("Failed to generate data: %v", err)
			}
			elapsed := time.Since(start)
			t.Logf("Data size: %d, Total time: %s ", size, elapsed)
		})

		t.Run(fmt.Sprintf("Read %d Records", size), func(t *testing.T) {
			start := time.Now()
			allTagID := tagRepo.GetAllTagIDs()

			for _, tagID := range allTagID {
				if err := cassandra.ExecuteQuery("Select * from tags where id = ?", tagID); err != nil {
					t.Errorf("Failed to execute query: %v", err)
				}
			}
			elapsed := time.Since(start)
			t.Logf("Read Operation : Data size: %d, Total time: %s ", size, elapsed)
		})

		t.Run(fmt.Sprintf("Update %d Records", size), func(t *testing.T) {
			start := time.Now()
			allTagID := tagRepo.GetAllTagIDs()
			for _, tagID := range allTagID {
				if err := cassandra.ExecuteQuery("Update tags SET properties = ? WHERE id = ?", utils.RandomString(100), tagID); err != nil {
					t.Errorf("Failed to execute query: %v", err)
				}
			}
			elapsed := time.Since(start)
			t.Logf("Update Operation : Data size: %d, Total time: %s ", size, elapsed)
		})

		t.Run(fmt.Sprintf("Update %d Records", size), func(t *testing.T) {
			start := time.Now()
			allTagID := tagRepo.GetAllTagIDs()
			for _, tagID := range allTagID {
				if err := cassandra.ExecuteQuery("Delete FROM tags WHERE id = ?", tagID); err != nil {
					t.Errorf("Failed to execute query: %v", err)
				}
			}
			elapsed := time.Since(start)
			t.Logf("Delete Operation : Data size: %d, Total time: %s ", size, elapsed)
		})
	}
}

func TestDBConcurrentOperationPerformance(t *testing.T) {
	//

	var wg *sync.WaitGroup = new(sync.WaitGroup)

	// Initialize the Cassandra connection
	cassandra, err := db.NewCassandra()
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}

	defer func() {
		// Truncate DB Tables
		fmt.Println(cassandra.TruncateTable([]string{"entities", "tags", "tag_entities"}))
		// Close the connection
		cassandra.Close()
	}()

	// Initialize repositories
	tagRepo := repositories.NewTagRepository(*cassandra)
	entityRepo := repositories.NewEntityRepository(*cassandra)
	tagEntityRepo := repositories.NewTagEntityRepository(*cassandra)

	// Truncate DB Tables
	fmt.Println(cassandra.TruncateTable([]string{"entities", "tags", "tag_entities"}))

	dataGenerator := &services.DataGeneratorService{
		TagRepo:       tagRepo,
		EntityRepo:    entityRepo,
		TagEntityRepo: tagEntityRepo,
	}

	dataSizes := []int{10}
	for _, size := range dataSizes {
		t.Run(fmt.Sprintf("%d Records", size), func(t *testing.T) {
			start := time.Now()
			if err := dataGenerator.GenerateData(*cassandra, size, false); err != nil {
				t.Errorf("Failed to generate data: %v", err)
			}
			elapsed := time.Since(start)
			t.Logf("Data size: %d, Total time: %s ", size, elapsed)
		})
		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Run(fmt.Sprintf("Read %d Records", size), func(t *testing.T) {
				start := time.Now()
				allTagID := tagRepo.GetAllTagIDs()

				for _, tagID := range allTagID {
					if err := cassandra.ExecuteQuery("Select * from tags where id = ?", tagID); err != nil {
						//t.Errorf("Failed to execute query: %v", err)
					}
				}
				elapsed := time.Since(start)
				t.Logf("Read Operation : Data size: %d, Total time: %s ", size, elapsed)
			})

		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Run(fmt.Sprintf("Update %d Records", size), func(t *testing.T) {
				start := time.Now()
				allTagID := tagRepo.GetAllTagIDs()
				for _, tagID := range allTagID {
					if err := cassandra.ExecuteQuery("Update tags SET properties = ? WHERE id = ?", utils.RandomString(100), tagID); err != nil {
						//t.Errorf("Failed to execute query: %v", err)
					}
				}
				elapsed := time.Since(start)
				t.Logf("Update Operation : Data size: %d, Total time: %s ", size, elapsed)
			})
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Run(fmt.Sprintf("Update %d Records", size), func(t *testing.T) {
				start := time.Now()
				allTagID := tagRepo.GetAllTagIDs()
				for _, tagID := range allTagID {
					if err := cassandra.ExecuteQuery("Delete FROM tags WHERE id = ?", tagID); err != nil {
						//t.Errorf("Failed to execute query: %v", err)
					}
				}
				elapsed := time.Since(start)
				t.Logf("Delete Operation : Data size: %d, Total time: %s ", size, elapsed)
			})
		}()

		wg.Wait()
	}
}

func TestDBIncerementalLoadSerialOperationPerformance(t *testing.T) {

	// Initialize the Cassandra connection
	cassandra, err := db.NewCassandra()
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}

	defer func() {
		// Truncate DB Tables
		fmt.Println(cassandra.TruncateTable([]string{"entities", "tags", "tag_entities"}))
		// Close the connection
		cassandra.Close()
	}()

	// Initialize repositories
	tagRepo := repositories.NewTagRepository(*cassandra)
	entityRepo := repositories.NewEntityRepository(*cassandra)
	tagEntityRepo := repositories.NewTagEntityRepository(*cassandra)

	// Truncate DB Tables
	fmt.Println(cassandra.TruncateTable([]string{"entities", "tags", "tag_entities"}))

	dataGenerator := &services.DataGeneratorService{
		TagRepo:       tagRepo,
		EntityRepo:    entityRepo,
		TagEntityRepo: tagEntityRepo,
	}

	dataSizes := []int{10, 10, 10, 10}
	for _, size := range dataSizes {
		t.Run(fmt.Sprintf("%d Records", size), func(t *testing.T) {
			start := time.Now()
			if err := dataGenerator.GenerateData(*cassandra, size, false); err != nil {
				t.Errorf("Failed to generate data: %v", err)
			}
			elapsed := time.Since(start)
			t.Logf("Data size: %d, Total time: %s ", size, elapsed)
		})

		t.Run(fmt.Sprintf("Read %d Records", size), func(t *testing.T) {
			start := time.Now()
			allTagID := tagRepo.GetAllTagIDs()

			for _, tagID := range allTagID {
				if err := cassandra.ExecuteQuery("Select * from tags where id = ?", tagID); err != nil {
					t.Errorf("Failed to execute query: %v", err)
				}
			}
			elapsed := time.Since(start)
			t.Logf("Read Operation : Data size: %d, Total time: %s ", size, elapsed)
		})

		t.Run(fmt.Sprintf("Update %d Records", size), func(t *testing.T) {
			start := time.Now()
			allTagID := tagRepo.GetAllTagIDs()
			for _, tagID := range allTagID {
				if err := cassandra.ExecuteQuery("Update tags SET properties = ? WHERE id = ?", utils.RandomString(100), tagID); err != nil {
					t.Errorf("Failed to execute query: %v", err)
				}
			}
			elapsed := time.Since(start)
			t.Logf("Update Operation : Data size: %d, Total time: %s ", size, elapsed)
		})

		t.Run(fmt.Sprintf("Update %d Records", size), func(t *testing.T) {
			start := time.Now()
			allTagID := tagRepo.GetAllTagIDs()
			for _, tagID := range allTagID {
				if err := cassandra.ExecuteQuery("Delete FROM tags WHERE id = ?", tagID); err != nil {
					t.Errorf("Failed to execute query: %v", err)
				}
			}
			elapsed := time.Since(start)
			t.Logf("Delete Operation : Data size: %d, Total time: %s ", size, elapsed)
		})
	}
}

func TestDBIncerementalLoadConcurrentOperationPerformance(t *testing.T) {

	var wg *sync.WaitGroup = new(sync.WaitGroup)
	// Initialize the Cassandra connection
	cassandra, err := db.NewCassandra()
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}

	defer func() {
		// Truncate DB Tables
		fmt.Println(cassandra.TruncateTable([]string{"entities", "tags", "tag_entities"}))
		// Close the connection
		cassandra.Close()
	}()

	// Initialize repositories
	tagRepo := repositories.NewTagRepository(*cassandra)
	entityRepo := repositories.NewEntityRepository(*cassandra)
	tagEntityRepo := repositories.NewTagEntityRepository(*cassandra)

	// Truncate DB Tables
	fmt.Println(cassandra.TruncateTable([]string{"entities", "tags", "tag_entities"}))

	dataGenerator := &services.DataGeneratorService{
		TagRepo:       tagRepo,
		EntityRepo:    entityRepo,
		TagEntityRepo: tagEntityRepo,
	}

	dataSizes := []int{10, 10, 10, 10, 10}
	for _, size := range dataSizes {
		t.Run(fmt.Sprintf("%d Records", size), func(t *testing.T) {
			start := time.Now()
			if err := dataGenerator.GenerateData(*cassandra, size, false); err != nil {
				t.Errorf("Failed to generate data: %v", err)
			}
			elapsed := time.Since(start)
			t.Logf("Data size: %d, Total time: %s ", size, elapsed)
		})
		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Run(fmt.Sprintf("Read %d Records", size), func(t *testing.T) {
				start := time.Now()
				allTagID := tagRepo.GetAllTagIDs()

				for _, tagID := range allTagID {
					if err := cassandra.ExecuteQuery("Select * from tags where id = ?", tagID); err != nil {
						//t.Errorf("Failed to execute query: %v", err)
					}
				}
				elapsed := time.Since(start)
				t.Logf("Read Operation : Data size: %d, Total time: %s ", size, elapsed)
			})

		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Run(fmt.Sprintf("Update %d Records", size), func(t *testing.T) {
				start := time.Now()
				allTagID := tagRepo.GetAllTagIDs()
				for _, tagID := range allTagID {
					if err := cassandra.ExecuteQuery("Update tags SET properties = ? WHERE id = ?", utils.RandomString(100), tagID); err != nil {
						//t.Errorf("Failed to execute query: %v", err)
					}
				}
				elapsed := time.Since(start)
				t.Logf("Update Operation : Data size: %d, Total time: %s ", size, elapsed)
			})
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			t.Run(fmt.Sprintf("Update %d Records", size), func(t *testing.T) {
				start := time.Now()
				allTagID := tagRepo.GetAllTagIDs()
				for _, tagID := range allTagID {
					if err := cassandra.ExecuteQuery("Delete FROM tags WHERE id = ?", tagID); err != nil {
						//t.Errorf("Failed to execute query: %v", err)
					}
				}
				elapsed := time.Since(start)
				t.Logf("Delete Operation : Data size: %d, Total time: %s ", size, elapsed)
			})
		}()

		wg.Wait()
	}
}
