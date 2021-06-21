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
	GetPostLocationById(id string, ctx context.Context) domain.PostLocation
	GetPostsByLocationName(location string, ctx context.Context) (*[]domain.PostLocation, error)
	SaveNewPostLocation(location domain.PostLocation, ctx context.Context) error
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

func (p postLocationRepo) GetPostLocationById(id string, ctx context.Context) domain.PostLocation {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var post domain.PostLocation
	err := p.collection.FindOne(ctx, bson.M{"post_id" : id}).Decode(&post)
	if err != nil {
		return domain.PostLocation{}
	}

	return post
}


func (p postLocationRepo) SaveNewPostLocation(location domain.PostLocation, ctx context.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := p.collection.InsertOne(ctx, location)
	if err != nil {
		return err
	}

	return nil
}


func (p postLocationRepo) GetPostsByLocationName(location string, ctx context.Context) (*[]domain.PostLocation, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	postsFiltered, err := p.collection.Find(ctx, bson.M{"location.location" : location})
	if err != nil {
		return nil, err
	}

	var posts []domain.PostLocation
	if err = postsFiltered.All(ctx, &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}


func NewPostLocationRepo(db *mongo.Client) PostLocationRepo {
	return &postLocationRepo {
		db : db,
		collection: db.Database("search_db").Collection("post_locations"),
	}
}