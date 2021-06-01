package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PostTagRepo interface {
	GetByPostId(id uint64) *mongo.SingleResult
	GetByTagId(id uint64) *mongo.SingleResult
}

type postTagRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (p postTagRepo) GetByTagId(id uint64) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := p.collection.FindOne(ctx, bson.M{"hashtag_id": id})
	return result
}

func (p postTagRepo) GetByPostId(id uint64) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := p.collection.FindOne(ctx, bson.M{"post_id": id})
	return result
}



func NewPostTagRepo(db *mongo.Client) PostTagRepo {
	return &postTagRepo{
		db: db,
		collection : db.Database("search_db").Collection("post_tags"),
	}
}