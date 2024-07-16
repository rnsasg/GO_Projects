package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/rnsasg/Tagging/config"
	"github.com/rnsasg/Tagging/controllers"
	"github.com/rnsasg/Tagging/repositories"
	"github.com/rnsasg/Tagging/routers"
	"github.com/rnsasg/Tagging/services"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting tagging service")
	config.LoadConfig()

	// Create a cassandra session
	cluster := gocql.NewCluster(config.GetEnv("CASSANDRA_HOST"))
	cluster.Keyspace = config.GetEnv("CASSANDRA_KEYSPACE")

	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatalf("Error creating Cassandra session: %v", err)
	}
	defer session.Close()

	entitiesRepository := repositories.NewEntityRepository(session)
	entityService := services.NewEntityService(entitiesRepository)
	entityController := controllers.NewEntityController(entityService)

	router := routers.InitTaggingRoutes(entityController)
	log.Fatal(http.ListenAndServe(":8080", router))
}
