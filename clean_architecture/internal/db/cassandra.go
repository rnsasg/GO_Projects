package db

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

type Cassandra struct {
	Session *gocql.Session
}

func NewCassandra() (*Cassandra, error) {
	cluster := gocql.NewCluster("127.0.0.1") // Update with actual Cassandra cluster IPs
	cluster.Keyspace = "tagging"
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 10 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Cassandra: %w", err)
	}

	if err := createSchema(session); err != nil {
		session.Close() // Ensure session is closed on error
		return nil, err
	}

	return &Cassandra{Session: session}, nil
}

func (c *Cassandra) Close() {
	if c.Session != nil {
		c.Session.Close()
	}
}

func createSchema(session *gocql.Session) error {
	if err := session.Query(`CREATE KEYSPACE IF NOT EXISTS tagging 
		WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1}`).Exec(); err != nil {
		return fmt.Errorf("failed to create keyspace: %w", err)
	}

	tableQueries := []string{
		`CREATE TABLE IF NOT EXISTS tags (
			id UUID PRIMARY KEY,
			owner_id INT,
			name TEXT,
			description TEXT,
			properties TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS entities (
			id UUID PRIMARY KEY,
			name TEXT,
			type TEXT,
			metadata TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS tag_entities (
			tag_id UUID,
			entity_id UUID,
			PRIMARY KEY (tag_id, entity_id)
		)`,
	}

	for _, query := range tableQueries {
		if err := session.Query(query).Exec(); err != nil {
			return fmt.Errorf("failed to create table with query [%s]: %w", query, err)
		}
	}

	log.Println("Schema created successfully.")
	return nil
}
