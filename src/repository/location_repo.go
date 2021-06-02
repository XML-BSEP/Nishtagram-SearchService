package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"search-service/domain"
	"time"
)

type locationRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}


type LocationRepo interface {
	ContainsLocation(location string, ctx context.Context) ([]domain.Location, error)
	ExactLocation(longitude float64, latitude float64, ctx context.Context) (domain.Location, error)
}


func (l locationRepo) ExactLocation(longitude float64, latitude float64, ctx context.Context) (domain.Location, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var  location domain.Location
	err := l.collection.FindOne(ctx, bson.M{"longitude" : longitude, "latitude" : latitude}).Decode(&location)
	if err != nil {
		return location, err
	}

	return location, nil



	/*
	var location domain.Location
	err := filterLocations.Decode(&location)

	if err != nil {
		log.Fatal(err)
	}

	return location, nil*/
}

func (l locationRepo) ContainsLocation(location string, ctx context.Context) ([]domain.Location, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterLocations, err := l.collection.Find(ctx, bson.M{"location" : primitive.Regex{Pattern: location, Options: "i"} })
	if err != nil {
		return nil, err
	}

	var locationsFiltered []domain.Location
	if err = filterLocations.All(ctx, &locationsFiltered); err != nil {
		return nil, err
	}

	return locationsFiltered, nil

}

func NewLocationRepo(db *mongo.Client) LocationRepo {
	return &locationRepo {
		db : db,
		collection: db.Database("search_db").Collection("locations"),
	}
}