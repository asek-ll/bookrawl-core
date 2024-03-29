package userbookstates

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status int

const (
	Empty Status = iota
	WantToRead
	Reading
	Readed
)

type State struct {
	UserId primitive.ObjectID `bson:"userId"`
	BookId primitive.ObjectID `bson:"bookId"`
	Status Status             `bson:"status,omitempty"`
	Rating int                `bson:"rating,omitempty"`
}
