package models

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"net/url"
)

type Entity struct {
	ID       gocql.UUID `json:"id"`
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Metadata string     `json:"metadata"`
}

func (e *Entity) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Entity) FromJSON(data []byte) error {
	return json.Unmarshal(data, e)
}

func (e *Entity) ToForm() url.Values {
	return url.Values{
		"id":       {e.ID.String()},
		"name":     {e.Name},
		"type":     {e.Type},
		"metadata": {e.Metadata},
	}
}
