package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/rnsasg/clean_architecture/config"
)

func main() {
	fmt.Println("Starting tagging service")
	//config.LoadConfig()

	// Create a cassandra session
	// cluster := gocql.NewCluster("cassandra")
	// cluster.Keyspace = "tagging"

	cluster := gocql.NewCluster(config.GetEnv("CASSANDRA_HOST"))
	cluster.Keyspace = config.GetEnv("CASSANDRA_KEYSPACE")
	cluster.Port = 9042
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Error creating Cassandra session: %v", err)
	}
	defer session.Close()

	log.Printf("Succes")

	// entitiesRepository := repositories.NewEntityRepository(session)
	// entityService := services.NewEntityService(entitiesRepository)
	// entityController := controllers.NewEntityController(entityService)

	// router := routers.InitTaggingRoutes(entityController)
	// log.Fatal(http.ListenAndServe(":8080", router))
}
