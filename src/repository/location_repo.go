package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type locationRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}


type LocationRepo interface {
	GetById(id uint64) *mongo.SingleResult
}


func (l locationRepo) GetById(id uint64) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	location := l.collection.FindOne(ctx, bson.M{"location_id" : id})
	return  location
}

func NewLocationRepo(db *mongo.Client) LocationRepo {
	return &locationRepo {
		db : db,
		collection: db.Database("search_db").Collection("locations"),
	}
}