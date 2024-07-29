package models

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"net/url"
)

// Tag represents a tag entity with db, gorm, and json struct tags.
type Tag struct {
	ID          gocql.UUID `db:"id" gorm:"column:id;primaryKey;type:uuid" json:"id"`
	OwnerID     int        `db:"owner_id" gorm:"column:owner_id;index" json:"owner_id"`
	Name        string     `db:"name" gorm:"column:name;type:text" json:"name"`
	Description string     `db:"description" gorm:"column:description;type:text" json:"description"`
	Properties  string     `db:"properties" gorm:"column:properties;type:text" json:"properties"`
}

// ToJSON serializes the Tag to JSON.
func (t *Tag) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

// FromJSON deserializes the Tag from JSON.
func (t *Tag) FromJSON(data []byte) error {
	return json.Unmarshal(data, t)
}

// ToForm converts the Tag to a URL-encoded form format.
func (t *Tag) ToForm() url.Values {
	return url.Values{
		"id":          {t.ID.String()},
		"owner_id":    {fmt.Sprint(t.OwnerID)},
		"name":        {t.Name},
		"description": {t.Description},
		"properties":  {t.Properties},
	}
}

// GetID returns the ID of the Tag.
func (t *Tag) GetID() gocql.UUID {
	return t.ID
}

// SetID sets the ID of the Tag.
func (t *Tag) SetID(id gocql.UUID) {
	t.ID = id
}

// GetOwnerID returns the OwnerID of the Tag.
func (t *Tag) GetOwnerID() int {
	return t.OwnerID
}

// SetOwnerID sets the OwnerID of the Tag.
func (t *Tag) SetOwnerID(ownerID int) {
	t.OwnerID = ownerID
}

// GetName returns the Name of the Tag.
func (t *Tag) GetName() string {
	return t.Name
}

// SetName sets the Name of the Tag.
func (t *Tag) SetName(name string) {
	t.Name = name
}

// GetDescription returns the Description of the Tag.
func (t *Tag) GetDescription() string {
	return t.Description
}

// SetDescription sets the Description of the Tag.
func (t *Tag) SetDescription(description string) {
	t.Description = description
}

// GetProperties returns the Properties of the Tag.
func (t *Tag) GetProperties() string {
	return t.Properties
}

// SetProperties sets the Properties of the Tag.
func (t *Tag) SetProperties(properties string) {
	t.Properties = properties
}
