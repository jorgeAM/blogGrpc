package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Blog models
type Blog struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title"`
	Content  string             `json:"content" bson:"content"`
	AuthorID string             `json:"authorId,omitempty" bson:"authorId"`
}
