package db

var tableQueries = []string{
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

const (
	keyspaceCreationQuery = `CREATE KEYSPACE IF NOT EXISTS tagging 
		WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1}`
)
