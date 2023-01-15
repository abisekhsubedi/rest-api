package models

// mongodb does not have a schema, so we can define things however we want but we need to be careful with that.
// The easiest way to be careful is to define a custom struct for each collection we want to use.
// serializeable fields

type Library struct {
	ID      string `json:"id" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Books  []Book `json:"books" bson:"books"`
}
