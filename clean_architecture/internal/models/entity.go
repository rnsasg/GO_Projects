package models

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"net/url"
)

// Entity represents an entity with a unique ID, name, type, and metadata.
type Entity struct {
	ID       gocql.UUID `json:"id" gorm:"primaryKey;column:id;type:uuid" db:"id"`
	Name     string     `json:"name" gorm:"column:name;type:text" db:"name"`
	Type     string     `json:"type" gorm:"column:type;type:text" db:"type"`
	Metadata string     `json:"metadata" gorm:"column:metadata;type:text" db:"metadata"`
}

// ToJSON serializes the Entity to JSON.
func (e *Entity) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// FromJSON deserializes the Entity from JSON.
func (e *Entity) FromJSON(data []byte) error {
	return json.Unmarshal(data, e)
}

// ToForm converts the Entity to a URL-encoded form format.
func (e *Entity) ToForm() url.Values {
	return url.Values{
		"id":       {e.ID.String()},
		"name":     {e.Name},
		"type":     {e.Type},
		"metadata": {e.Metadata},
	}
}

// GetID returns the ID of the Entity.
func (e *Entity) GetID() gocql.UUID {
	return e.ID
}

// SetID sets the ID of the Entity.
func (e *Entity) SetID(id gocql.UUID) {
	e.ID = id
}

// GetName returns the Name of the Entity.
func (e *Entity) GetName() string {
	return e.Name
}

// SetName sets the Name of the Entity.
func (e *Entity) SetName(name string) {
	e.Name = name
}

// GetType returns the Type of the Entity.
func (e *Entity) GetType() string {
	return e.Type
}

// SetType sets the Type of the Entity.
func (e *Entity) SetType(etype string) {
	e.Type = etype
}

// GetMetadata returns the Metadata of the Entity.
func (e *Entity) GetMetadata() string {
	return e.Metadata
}

// SetMetadata sets the Metadata of the Entity.
func (e *Entity) SetMetadata(metadata string) {
	e.Metadata = metadata
}
