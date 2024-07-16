package models

type Tag struct {
	ID          string // Index for Storing Tags
	OwnerID     uint8  // XIQ Owner ID
	Name        string // Name of the Tag
	Description string // Description of the Tag
	Properties  string // Properties of the Tag
	EnityList   []Entity
}
