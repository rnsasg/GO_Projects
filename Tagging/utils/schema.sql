CREATE KEYSPACE tagging WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
USE tagging;


CREATE TABLE IF NOT EXISTS tags (
    id UUID PRIMARY KEY,
    name TEXT,
    parent UUID,
)

CREATE TABLE IF NOT EXISTS entities (
    id UUID PRIMARY KEY,
    name TEXT,
    type TEXT,
)

CREATE TABLE IF NOT EXISTS attributes (
    attribute_id UUID,
    key text,
    value text,
)

CREATE TABLE IF NOT EXISTS tag_entities (
    tag_id UUID,
    entity_id UUID,
    PRIMARY KEY (tag_id, entity_id)
)

CREATE TABLE IF NOT EXISTS tag_attributes (
    tag_id UUID,
    attribute_id UUID,
    PRIMARY KEY (tag_id, attribute_id)
)



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

