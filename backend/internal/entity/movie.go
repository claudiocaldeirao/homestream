package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name"`
	Path     string             `bson:"path"`
	Ext      string             `bson:"ext"`
	Scanned  int64              `bson:"scanned_at"`
	Metadata OmdbMetadata
}
