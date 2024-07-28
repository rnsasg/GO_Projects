package models

import (
	"encoding/json"
	"net/url"

	"github.com/gocql/gocql"
)

type TagEntity struct {
	TagID    gocql.UUID `json:"tag_id"`
	EntityID gocql.UUID `json:"entity_id"`
}

func (te *TagEntity) ToJSON() ([]byte, error) {
	return json.Marshal(te)
}

func (te *TagEntity) FromJSON(data []byte) error {
	return json.Unmarshal(data, te)
}

func (te *TagEntity) ToForm() url.Values {
	return url.Values{
		"tag_id":    {te.TagID.String()},
		"entity_id": {te.EntityID.String()},
	}
}
