package authors

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Collection *mongo.Collection
}

func (store *Store) Upsert(author *Author) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: author.Id}}
	update := bson.D{{Key: "$set", Value: author}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := store.Collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if result.UpsertedID != nil {
		author.Id = result.UpsertedID.(primitive.ObjectID)
	}

	return nil
}

func (store *Store) UpsertManyByFantlabId(authors []*Author) error {
	models := make([]mongo.WriteModel, len(authors))

	for i, author := range authors {
		filter := bson.D{{Key: "fantlabId", Value: author.FantlabId}}
		update := bson.D{{Key: "$set", Value: author}}
		models[i] = mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
	}

	opts := options.BulkWrite().SetOrdered(false)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := store.Collection.BulkWrite(ctx, models, opts)
	return err
}

func isAuthorMatchedStrict(name string, author *Author) bool {
	if author == nil {
		return false
	}

	if isNameMatchedStrict(name, author.Name) {
		return true
	}

	for _, alias := range author.Aliases {
		if isNameMatchedStrict(name, alias) {
			return true
		}
	}
	return false
}

func isNameMatchedStrict(name string, authorName string) bool {
	name = strings.ReplaceAll(strings.ToLower(name), "ё", "е")
	authorName = strings.ReplaceAll(strings.ToLower(authorName), "ё", "е")

	authorsWords := make(map[string]bool)
	for _, s := range strings.Split(authorName, " ") {
		authorsWords[strings.ToLower(s)] = true
	}

	for _, s := range strings.Split(name, " ") {
		if _, e := authorsWords[strings.ToLower(s)]; !e {
			fmt.Println("compare", name, "==", authorName, false)
			return false
		}
	}

	fmt.Println("compare", name, "==", authorName, true)

	return true
}

func (store *Store) FindByName(name string) (*Author, error) {
	filter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: name}}}}
	sort := bson.D{{Key: "score", Value: bson.D{{Key: "$meta", Value: "textScore"}}}}

	opts := options.Find().SetSort(sort).SetLimit(3)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := store.Collection.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var models []Author
	err = cursor.All(ctx, &models)

	if err != nil {
		return nil, err
	}

	for _, author := range models {
		if isAuthorMatchedStrict(name, &author) {
			return &author, nil
		}
	}

	return nil, nil
}

func (store *Store) findOne(filter interface{}) (*Author, error) {

	opts := &options.FindOneOptions{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := store.Collection.FindOne(ctx, filter, opts)

	if result.Err() != nil {
		return nil, result.Err()
	}

	var author Author
	err := result.Decode(&author)
	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (store *Store) findMany(filter interface{}) ([]Author, error) {

	opts := options.Find()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := store.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var models []Author
	err = result.All(ctx, &models)

	if err != nil {
		return nil, err
	}

	return models, nil
}

func (store *Store) FindOneById(id primitive.ObjectID) (*Author, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	return store.findOne(filter)
}

func (store *Store) FindManyByFantlabId(fantlabIds []int) ([]Author, error) {
	filter := bson.D{{Key: "fantlabId", Value: bson.D{{Key: "$in", Value: fantlabIds}}}}
	return store.findMany(filter)
}
