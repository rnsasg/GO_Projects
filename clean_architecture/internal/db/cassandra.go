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

	return &Cassandra{Session: session}, nil
}

func (c *Cassandra) Close() {
	if c.Session != nil {
		c.Session.Close()
	}
}

func (c *Cassandra) createSchema() error {

	if err := c.ExecuteQuery(keyspaceCreationQuery); err != nil {
		c.Session.Close()
		return fmt.Errorf("failed to create keyspace: %w", err)
	}

	for _, query := range tableQueries {
		if err := c.ExecuteQuery(query); err != nil {
			// Ensure session is closed on error
			c.Session.Close()
			return fmt.Errorf("failed to create table with query [%s]: %w", query, err)
		}
	}

	log.Println("Schema created successfully.")
	return nil
}

func (c *Cassandra) TruncateTable(tables []string) error {
	for _, table := range tables {
		if err := c.ExecuteQuery(fmt.Sprintf("truncate %s", table)); err != nil {
			return fmt.Errorf("failed to truncate table: %s , %w", table, err)
		}
	}
	return nil
}

func (c *Cassandra) DropTable(tables []string) error {
	for _, table := range tables {
		if err := c.ExecuteQuery(fmt.Sprintf("truncate %s", table)); err != nil {
			return fmt.Errorf("failed to truncate table: %s , %w", table, err)
		}
	}
	return nil
}

func (c *Cassandra) ExecuteQuery(query string, values ...interface{}) error {
	if err := c.Session.Query(query, values...).Bind(values...).Exec(); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (c *Cassandra) Iterator(query string, values ...interface{}) *gocql.Iter {
	return c.Session.Query(query, values...).Bind(values...).Iter()
}
