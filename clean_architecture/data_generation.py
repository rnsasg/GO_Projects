from cassandra.cluster import Cluster
from uuid import uuid4
import random
import string

# Connect to the Cassandra cluster
cluster = Cluster(['localhost'])
session = cluster.connect()

# Create keyspace and use it
session.execute("""
CREATE KEYSPACE IF NOT EXISTS tagging 
WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
""")
session.set_keyspace('tagging')

# Create tables
session.execute("""
CREATE TABLE IF NOT EXISTS tags (
    id UUID PRIMARY KEY,
    owner_id INT,
    name TEXT,
    description TEXT,
    properties TEXT
);
""")

session.execute("""
CREATE TABLE IF NOT EXISTS entities (
    id UUID PRIMARY KEY,
    name TEXT,
    type TEXT,
    metadata TEXT
);
""")

session.execute("""
CREATE TABLE IF NOT EXISTS tag_entities (
    tag_id UUID,
    entity_id UUID,
    PRIMARY KEY (tag_id, entity_id)
);
""")

# Generate random data
def random_string(length=10):
    letters = string.ascii_letters
    return ''.join(random.choice(letters) for i in range(length))

# Insert records into the 'tags' table
tags = []
for _ in range(1000000):
    tag_id = uuid4()
    tags.append(tag_id)
    session.execute(
        """
        INSERT INTO tags (id, owner_id, name, description, properties)
        VALUES (%s, %s, %s, %s, %s)
        """,
        (tag_id, random.randint(1, 1000), random_string(15), random_string(50), random_string(100))
    )

print("Inserted 1 million records into 'tags'")

# Insert records into the 'entities' table
entities = []
for _ in range(1000000):
    entity_id = uuid4()
    entities.append(entity_id)
    session.execute(
        """
        INSERT INTO entities (id, name, type, metadata)
        VALUES (%s, %s, %s, %s)
        """,
        (entity_id, random_string(15), random_string(10), random_string(100))
    )

print("Inserted 1 million records into 'entities'")

# Insert records into the 'tag_entities' table using IDs from 'tags' and 'entities'
batch_size = 1000  # Number of records to insert in a batch
for i in range(0, 1000000, batch_size):
    batch = []
    for j in range(batch_size):
        tag_id = tags[random.randint(0, len(tags) - 1)]
        entity_id = entities[random.randint(0, len(entities) - 1)]
        batch.append((tag_id, entity_id))
    prepared = session.prepare("INSERT INTO tag_entities (tag_id, entity_id) VALUES (?, ?)")
    for record in batch:
        session.execute(prepared, record)
    print(f"Inserted {i + batch_size} records into 'tag_entities'")

# Close the connection
cluster.shutdown()
