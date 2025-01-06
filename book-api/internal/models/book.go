package models

type Book struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Title    string `json:"title" bson:"title"`
	Category string `json:"category" bson:"category"`
	Author   string `json:"author" bson:"author"`
	Synopsis string `json:"synopsis" bson:"synopsis"`
}
