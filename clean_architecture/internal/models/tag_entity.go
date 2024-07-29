package models

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"net/url"
)

// TagEntity represents the association between a tag and an entity.
type TagEntity struct {
	TagID    gocql.UUID `db:"tag_id" gorm:"column:tag_id;primaryKey;type:uuid" json:"tag_id"`
	EntityID gocql.UUID `db:"entity_id" gorm:"column:entity_id;primaryKey;type:uuid" json:"entity_id"`
}

// ToJSON serializes the TagEntity to JSON.
func (te *TagEntity) ToJSON() ([]byte, error) {
	return json.Marshal(te)
}

// FromJSON deserializes the TagEntity from JSON.
func (te *TagEntity) FromJSON(data []byte) error {
	return json.Unmarshal(data, te)
}

// ToForm converts the TagEntity to a URL-encoded form format.
func (te *TagEntity) ToForm() url.Values {
	return url.Values{
		"tag_id":    {te.TagID.String()},
		"entity_id": {te.EntityID.String()},
	}
}

// GetTagID returns the TagID of the TagEntity.
func (te *TagEntity) GetTagID() gocql.UUID {
	return te.TagID
}

// SetTagID sets the TagID of the TagEntity.
func (te *TagEntity) SetTagID(tagID gocql.UUID) {
	te.TagID = tagID
}

// GetEntityID returns the EntityID of the TagEntity.
func (te *TagEntity) GetEntityID() gocql.UUID {
	return te.EntityID
}

// SetEntityID sets the EntityID of the TagEntity.
func (te *TagEntity) SetEntityID(entityID gocql.UUID) {
	te.EntityID = entityID
}
