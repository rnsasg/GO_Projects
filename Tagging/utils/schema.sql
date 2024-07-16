CREATE KEYSPACE tagging WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE tagging;

CREATE TABLE entities (
    id UUID PRIMARY KEY,
    name TEXT,
    type TEXT,
    metadata TEXT
);