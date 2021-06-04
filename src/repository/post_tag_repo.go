package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"search-service/domain"
	"time"
)

type PostTagRepo interface {
	GetPostsByHashTag(hashTag string, ctx context.Context) ([]string, error)
}

type postTagRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}

func (p postTagRepo) GetPostsByHashTag(hashTag string, ctx context.Context) ([]string, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterPostst, err := p.collection.Find(ctx, bson.M{"hashtag" : hashTag })

	var postsFiltered []domain.PostLocation
	if err = filterPostst.All(ctx, &postsFiltered); err != nil {
		return nil, err
	}

	var sliceIds []string
	for _, p := range postsFiltered {
		sliceIds = append(sliceIds, p.PostId)
	}
	return sliceIds, nil
}

func NewPostTagRepo(db *mongo.Client) PostTagRepo {
	return &postTagRepo{
		db: db,
		collection : db.Database("search_db").Collection("post_tags"),
	}
}