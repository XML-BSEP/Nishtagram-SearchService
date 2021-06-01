package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type postLocationRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

type PostLocationRepo interface {
	GetById(id uint64) *mongo.SingleResult
}

func (l postLocationRepo) GetById(id uint64) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	postLocation := l.collection.FindOne(ctx, bson.M{"post_id" : id})
	return  postLocation
}


func NewPostLocationRepo(db *mongo.Client) PostLocationRepo {
	return &postLocationRepo {
		db : db,
		collection: db.Database("search_db").Collection("post_locations"),
	}
}