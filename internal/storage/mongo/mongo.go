package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/seggga/approve-auth/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database ...
type Database struct {
	client *mongo.Client
	coll   *mongo.Collection
}

// New creates a mongodb client and creates a database with given data
func New(connString string, users map[string]entity.UserOpts) (*Database, error) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(connString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to mongo instance %s: %w", connString, err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error pinging mongo instance %s: %w", connString, err)
	}

	// create database -> create collection -> add data
	collection := client.Database("team9").Collection("users-coll")
	data := make([]interface{}, 0, len(users))
	for _, v := range users {
		data = append(data, v)
	}
	_, err = collection.InsertMany(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("error adding data to mongo collection: %w", err)
	}

	return &Database{
		client: client,
		coll:   collection,
	}, nil
}

// ReadUser extracts a user by login
func (d *Database) ReadUser(login string) (*entity.UserOpts, error) {
	// filter := bson.D
	filter := bson.D{
		primitive.E{Key: "login", Value: login},
	}
	// TODO: move context to method signature
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	user := entity.UserOpts{}
	err := d.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("user with login %s was not found", login)
	}

	return &user, nil
}
