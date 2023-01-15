package userbookstates

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status int

const (
	Empty Status = iota
	WantToRead
	Reading
	Readed
)

type BookRead struct {
	AbookId string    `bson:"abookId,omitempty"`
	Start   time.Time `bson:"start,omitempty"`
	End     time.Time `bson:"end,omitempty"`
}

type State struct {
	UserId primitive.ObjectID `bson:"userId"`
	BookId primitive.ObjectID `bson:"bookId"`
	Status Status             `bson:"status,omitempty"`
	Rating int                `bson:"rating,omitempty"`
	Reads  []BookRead         `bson:"reads,omitempty"`
}
