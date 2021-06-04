package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"search-service/domain"
	"time"
)

type postLocationRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}



type PostLocationRepo interface {
	GetPostsByExactLocation(longitude float64, latitude float64 , ctx context.Context) ([]string, error)
	GetPostsByLocationContains(location string, ctx context.Context) ([]string, error)
}

func (p postLocationRepo) GetPostsByLocationContains(location string, ctx context.Context) ([]string, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterPosts, err := p.collection.Find(ctx, bson.M{"location.location" : primitive.Regex{Pattern: location, Options: "i"} })
	if err != nil {
		return nil, err
	}

	var postsFiltered []domain.PostLocation
	if err = filterPosts.All(ctx, &postsFiltered); err != nil {
		return nil, err
	}

	var sliceIds []string
	for _, p := range postsFiltered {
		sliceIds = append(sliceIds, p.PostId)
	}
	return sliceIds, nil
}

func (p postLocationRepo) GetPostsByExactLocation(longitude float64, latitude float64 , ctx context.Context) ([]string, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterPosts, err := p.collection.Find(ctx, bson.M{"location.longitude" : longitude, "location.latitude" : latitude})
	if err != nil {
		return nil, err
	}

	var postsFiltered []domain.PostLocation
	if err = filterPosts.All(ctx, &postsFiltered); err != nil {
		return nil, err
	}

	var sliceIds []string
	for _, p := range postsFiltered {
		sliceIds = append(sliceIds, p.PostId)
	}
	return sliceIds, nil
}

func NewPostLocationRepo(db *mongo.Client) PostLocationRepo {
	return &postLocationRepo {
		db : db,
		collection: db.Database("search_db").Collection("post_locations"),
	}
}