package books

import "go.mongodb.org/mongo-driver/bson/primitive"

type FantLabRating struct {
	Rating float32 `bson:"rating"`
	Voters int     `bson:"voters"`
}

type Book struct {
	Id            primitive.ObjectID   `bson:"_id,omitempty"`
	Name          string               `bson:"name"`
	Authors       []primitive.ObjectID `bson:"authors"`
	FantLabId     *int                 `bson:"fantlabId,omitempty"`
	FantLabRating *FantLabRating       `bson:"fantlabRating,omitempty"`
}
