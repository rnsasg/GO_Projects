package models

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gocql/gocql"
)

type Tag struct {
	ID          gocql.UUID `json:"id"`
	OwnerID     int        `json:"owner_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Properties  string     `json:"properties"`
}

func (t *Tag) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Tag) FromJSON(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t *Tag) ToForm() url.Values {
	return url.Values{
		"id":          {t.ID.String()},
		"owner_id":    {fmt.Sprint(t.OwnerID)},
		"name":        {t.Name},
		"description": {t.Description},
		"properties":  {t.Properties},
	}
}
