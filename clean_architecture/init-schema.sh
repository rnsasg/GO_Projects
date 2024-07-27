#!/bin/bash

# Maximum number of attempts to wait for Cassandra to start
MAX_ATTEMPTS=20
ATTEMPT=1

echo "Waiting for Cassandra to start..."

while [ $ATTEMPT -le $MAX_ATTEMPTS ]; do
  # Attempt to run a simple cqlsh command
  if cqlsh -e "SELECT now() FROM system.local" > /dev/null 2>&1; then
    echo "Cassandra is up and running."
    break
  else
    echo "Cassandra is unavailable - retrying in 5 seconds... (Attempt: $ATTEMPT/$MAX_ATTEMPTS)"
    ATTEMPT=$((ATTEMPT + 1))
    sleep 5
  fi
done

if [ $ATTEMPT -gt $MAX_ATTEMPTS ]; then
  echo "Cassandra did not start in time. Exiting."
  exit 1
fi

# Run your CQL commands to set up the schema
echo "Setting up keyspace and tables..."
cqlsh -e "
CREATE KEYSPACE IF NOT EXISTS tagging WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
USE tagging;

CREATE TABLE IF NOT EXISTS tags (
    id UUID PRIMARY KEY,
    owner_id INT,
    name TEXT,
    description TEXT,
    properties TEXT
);

CREATE TABLE IF NOT EXISTS entities (
    id UUID PRIMARY KEY,
    name TEXT,
    type TEXT,
    metadata TEXT
);

CREATE TABLE IF NOT EXISTS tag_entities (
    tag_id UUID,
    entity_id UUID,
    PRIMARY KEY (tag_id, entity_id)
);
"

echo "Schema setup completed."
