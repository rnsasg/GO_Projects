package utils

import (
	"github.com/gocql/gocql"
	"github.com/rnsasg/clean_architecture/internal/db"
)

func (c *db.Cassandra) QueryTags() {
	iter := c.Iterator("select * from tags")
	defer iter.Close()
}

func (c *db.Cassandra QueryEntities(session *gocql.Session) {
	iter := c.Iterator(`SELECT * FROM entities`)
	defer iter.Close()
}
