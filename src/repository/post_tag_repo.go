package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type PostTagRepo interface {
	GetByPostId(id string) *mongo.SingleResult
	ContainsHashTag(hashTag string) *mongo.SingleResult
}

type postTagRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (p postTagRepo) GetByPostId(id string) *mongo.SingleResult {
	panic("implement me")
}

func (p postTagRepo) ContainsHashTag(hashTag string) *mongo.SingleResult {
	panic("implement me")
}

func NewPostTagRepo(db *mongo.Client) PostTagRepo {
	return &postTagRepo{
		db: db,
		collection : db.Database("search_db").Collection("post_tags"),
	}
}