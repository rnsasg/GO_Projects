package models

import "github.com/gocql/gocql"

type Entity struct {
	ID       gocql.UUID `json:"id"`
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Metadata string     `json:"metadata"`
}
