package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"time"
)

func randomString(length int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, length)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}

func main() {
	// Set up the cluster configuration
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "tagging"
	cluster.Consistency = gocql.Quorum

	// Connect to the cluster
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Create keyspace and use it
	if err := session.Query(`CREATE KEYSPACE IF NOT EXISTS tagging WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1}`).Exec(); err != nil {
		log.Fatal(err)
	}

	// Create tables
	if err := session.Query(`CREATE TABLE IF NOT EXISTS tags (
		id UUID PRIMARY KEY,
		owner_id INT,
		name TEXT,
		description TEXT,
		properties TEXT
	)`).Exec(); err != nil {
		log.Fatal(err)
	}

	if err := session.Query(`CREATE TABLE IF NOT EXISTS entities (
		id UUID PRIMARY KEY,
		name TEXT,
		type TEXT,
		metadata TEXT
	)`).Exec(); err != nil {
		log.Fatal(err)
	}

	if err := session.Query(`CREATE TABLE IF NOT EXISTS tag_entities (
		tag_id UUID,
		entity_id UUID,
		PRIMARY KEY (tag_id, entity_id)
	)`).Exec(); err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

	// Generate and insert 1 million records into 'tags' table
	for i := 0; i < 10; i++ {
		tagID, _ := gocql.ParseUUID(uuid.New().String())
		ownerID := rand.Intn(1000) + 1
		name := randomString(15)
		description := randomString(50)
		properties := randomString(100)
		if err := session.Query(`INSERT INTO tags (id, owner_id, name, description, properties) VALUES (?, ?, ?, ?, ?)`,
			tagID, ownerID, name, description, properties).Exec(); err != nil {
			log.Fatal(err)
		}
		if i%1000 == 0 {
			fmt.Printf("Inserted %d records into 'tags'\n", i+1)
		}
	}

	// Generate and insert 1 million records into 'entities' table
	for i := 0; i < 10; i++ {
		entityID, _ := gocql.ParseUUID(uuid.New().String())
		name := randomString(15)
		entityType := randomString(10)
		metadata := randomString(100)
		if err := session.Query(`INSERT INTO entities (id, name, type, metadata) VALUES (?, ?, ?, ?)`,
			entityID, name, entityType, metadata).Exec(); err != nil {
			log.Fatal(err)
		}
		if i%1000 == 0 {
			fmt.Printf("Inserted %d records into 'entities'\n", i+1)
		}
	}

	// Select all tag IDs and entity IDs to use in the tag_entities table
	tagIDs := []gocql.UUID{}
	entityIDs := []gocql.UUID{}

	iter := session.Query(`SELECT id FROM tags`).Iter()
	var id gocql.UUID
	for iter.Scan(&id) {
		id, _ := gocql.ParseUUID(uuid.New().String())
		tagIDs = append(tagIDs, id)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	iter = session.Query(`SELECT id FROM entities`).Iter()
	for iter.Scan(&id) {
		id, _ := gocql.ParseUUID(uuid.New().String())
		entityIDs = append(entityIDs, id)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	// Generate and insert records into 'tag_entities' table using IDs from 'tags' and 'entities'
	for i := 0; i < 10; i++ {
		//entityID,_:= gocql.ParseUUID(uuid.New().String())
		tagID, _ := gocql.ParseUUID(tagIDs[rand.Intn(len(tagIDs))].String())
		entityID, _ := gocql.ParseUUID(entityIDs[rand.Intn(len(entityIDs))].String())

		if err := session.Query(`INSERT INTO tag_entities (tag_id, entity_id) VALUES (?, ?)`,
			tagID, entityID).Exec(); err != nil {
			log.Fatal(err)
		}
		if i%1000 == 0 {
			fmt.Printf("Inserted %d records into 'tag_entities'\n", i+1)
		}
	}

	fmt.Println("Data generation completed.")
}
