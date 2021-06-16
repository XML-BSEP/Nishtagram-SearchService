package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"search-service/domain"
	"time"
)

type PostTagRepo interface {
	GetPostsByHashTag(hashTag string, ctx context.Context) ([]string, error)
	GetPostTagById(id string, ctx context.Context) domain.PostTag
	GetPostsBbyHashTagName(hashtag string, ctx context.Context) (*[]domain.PostTag, error)
}

type postTagRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}



func (p postTagRepo) GetPostsByHashTag(hashTag string, ctx context.Context) ([]string, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterPostst, err := p.collection.Find(ctx, bson.M{"hashtag" : primitive.Regex{Pattern: hashTag, Options: "i"} })

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


func (p postTagRepo) GetPostTagById(id string, ctx context.Context) domain.PostTag {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var post domain.PostTag
	err := p.collection.FindOne(ctx, bson.M{"post_id" : id}).Decode(&post)
	if err != nil {
		return domain.PostTag{}
	}

	return post
}


func (p postTagRepo) GetPostsBbyHashTagName(hashtag string, ctx context.Context) (*[]domain.PostTag, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	postsFiltered, err := p.collection.Find(ctx, bson.M{"hashtag" : hashtag})
	if err != nil {
		return nil, err
	}

	var posts []domain.PostTag
	if err = postsFiltered.All(ctx, &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}


func NewPostTagRepo(db *mongo.Client) PostTagRepo {
	return &postTagRepo{
		db: db,
		collection : db.Database("search_db").Collection("post_tags"),
	}
}