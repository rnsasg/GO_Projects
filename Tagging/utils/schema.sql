CREATE KEYSPACE tagging WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
USE tagging;

CREATE TABLE tags (
    id UUID PRIMARY KEY,
    owner_id INT,
    name TEXT,
    description TEXT,
    properties TEXT
);

CREATE TABLE entities (
    id UUID PRIMARY KEY,
    name TEXT,
    type TEXT,
    metadata TEXT
);

CREATE TABLE tag_entities (
    tag_id UUID,
    entity_id UUID,
    PRIMARY KEY (tag_id, entity_id)
);