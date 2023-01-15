package models

type Book struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Title  string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
	ISBN   string `json:"isbn" bson:"isbn"`
}
