package authors

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FantlabId int                `bson:"fantlabId,omitempty"`
	Name      string             `bson:"name"`
	Aliases   []string           `bson:"aliases"`
}
