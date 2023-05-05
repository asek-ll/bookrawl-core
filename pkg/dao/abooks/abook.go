package abooks

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ABook struct {
	Id          string
	RawTitle    string
	Title       string
	Author      string
	Artists     []string
	Year        int
	Date        time.Time
	Link        string
	Description string
	Length      int
	Size        string
	Quality     string
	Props       map[string]string
	AuthorId    []int
	Authors     []primitive.ObjectID `bson:"authors,omitempty"`
}
