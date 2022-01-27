package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FileTag holds the schema definition for the FileTag entity.
type FileTag struct {
	ent.Schema
}

// Fields of the FileTag.
func (FileTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("from_file"),
	}
}

// Edges of the FileTag.
func (FileTag) Edges() []ent.Edge {
	// references the tag and file
	return nil
}
