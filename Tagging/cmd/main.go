package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/rnsasg/Tagging/controllers"
	"github.com/rnsasg/Tagging/repositories"
	"github.com/rnsasg/Tagging/routers"
	"github.com/rnsasg/Tagging/services"
)

func main() {
	fmt.Println("Starting tagging service")
	//config.LoadConfig()

	// Create a cassandra session
	cluster := gocql.NewCluster("172.17.0.2:9042")
	cluster.Keyspace = "tagging"

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Error creating Cassandra session: %v", err)
	}
	defer session.Close()

	log.Printf("Succes")

	entitiesRepository := repositories.NewEntityRepository(session)
	entityService := services.NewEntityService(entitiesRepository)
	entityController := controllers.NewEntityController(entityService)

	router := routers.InitTaggingRoutes(entityController)
	log.Fatal(http.ListenAndServe(":8080", router))
}
