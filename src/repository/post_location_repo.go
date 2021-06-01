package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type postLocationRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}


type PostLocationRepo interface {
	GetByPostId(id string, ctx context.Context) *mongo.SingleResult
}


func (p postLocationRepo) GetByPostId(id string, ctx context.Context) *mongo.SingleResult {
	panic("implement me")
}

func NewPostLocationRepo(db *mongo.Client) PostLocationRepo {
	return &postLocationRepo {
		db : db,
		collection: db.Database("search_db").Collection("post_locations"),
	}
}