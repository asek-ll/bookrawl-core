package users

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id              primitive.ObjectID   `bson:"_id,omitempty"`
	ChatId          *int64               `bson:"chatId,omitempty"`
	FavoriteAuthors []primitive.ObjectID `bson:"favoriteAuthors,omitempty"`
}

type Store struct {
	Collection *mongo.Collection
}

func (s *Store) Create(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := s.Collection.InsertOne(ctx, user)
	return err
}

func (s *Store) FindById(id primitive.ObjectID) (*User, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := s.Collection.FindOne(ctx, filter)

	var user User
	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (s *Store) Upsert(user *User) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: user.Id}}
	update := bson.D{{Key: "$set", Value: user}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := s.Collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (s *Store) FindByFavoriteAuthors(favoriteAuthorIds []int) ([]User, error) {
	filter := bson.D{{Key: "favoriteAuthors", Value: bson.D{{Key: "$in", Value: favoriteAuthorIds}}}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := s.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var models []User
	err = cursor.All(ctx, &models)

	if err != nil {
		return nil, err
	}

	return models, nil
}

func (s *Store) GetOrCreateByChatId(id int64) (*User, error) {
	filter := bson.D{{Key: "chatId", Value: id}}

	opts := &options.FindOneAndUpdateOptions{}
	opts.SetUpsert(true)

	update := bson.D{{Key: "$setOnInsert", Value: bson.D{
		{Key: "chatId", Value: id},
	}}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := s.Collection.FindOneAndUpdate(ctx, filter, update, opts)

	if result.Err() != nil {
		return nil, result.Err()
	}

	var user User
	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Store) updateOne(userId primitive.ObjectID, update interface{}) (*User, error) {
	filter := bson.D{{Key: "_id", Value: userId}}

	opts := &options.FindOneAndUpdateOptions{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := s.Collection.FindOneAndUpdate(ctx, filter, update, opts)

	if result.Err() != nil {
		return nil, result.Err()
	}

	var user User
	err := result.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Store) AddFavoriteAuthor(userId primitive.ObjectID, authorId primitive.ObjectID) (*User, error) {
	update := bson.D{{Key: "$addToSet", Value: bson.D{
		{Key: "favoriteAuthors", Value: authorId},
	}}}

	return s.updateOne(userId, update)
}
