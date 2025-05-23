package entity

type Movie struct {
	Name     string `bson:"name"`
	Path     string `bson:"path"`
	Ext      string `bson:"ext"`
	Scanned  int64  `bson:"scanned_at"`
	Metadata OmdbMetadata
}
