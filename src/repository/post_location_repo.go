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
	Create(postLocation domain.PostLocation, ctx context.Context) error
	GetById(id string, location string, ctx context.Context) (*domain.PostLocation, error)
}

func (p *postLocationRepo) GetPostsByLocationContains(location string, ctx context.Context) ([]string, error) {
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

func (p *postLocationRepo) Create(postLocation domain.PostLocation, ctx context.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	value, err := p.GetById(postLocation.PostId, postLocation.Location.LocationName, ctx)

	if err == nil {
		filter := bson.M{"_id": value.ID}
		_, err := p.collection.ReplaceOne(ctx, filter, postLocation)
		if err != nil {
			return err
		}
		return nil
	}
	newObj := bson.D{
		{"post_id", postLocation.PostId},
		{"user_id", postLocation.UserId},
		{"location", domain.Location{
			LocationName: postLocation.Location.LocationName,
			Longitude: postLocation.Location.Longitude,
			Latitude:  postLocation.Location.Latitude,
		}},

	}
	_, err = p.collection.InsertOne(ctx, newObj)
	return err

}


func (p *postLocationRepo) GetById(id string, location string, ctx context.Context) (*domain.PostLocation, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var postLocation *domain.PostLocation
	err := p.collection.FindOne(ctx, bson.M{"post_id" : id}).Decode(&postLocation)
	if err != nil {
		return postLocation, err
	}

	return postLocation, nil
}

func NewPostLocationRepo(db *mongo.Client) PostLocationRepo {
	return &postLocationRepo {
		db : db,
		collection: db.Database("search_db").Collection("post_locations"),
	}
}