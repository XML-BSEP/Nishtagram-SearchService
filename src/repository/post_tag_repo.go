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
	Create(postTag domain.PostTags, ctx context.Context) error
	GetPostTagsById(id string, ctx context.Context) (*domain.PostTags, error)
	GetByIdAndHashTag(id string, hashTag string, ctx context.Context) (*domain.PostTags, error)
}

type postTagRepo struct {
	collection *mongo.Collection
	db *mongo.Client
}


func (p *postTagRepo) Create(postTag domain.PostTags, ctx context.Context) error {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	value, err := p.GetByIdAndHashTag(postTag.PostId, postTag.Hashtag, ctx)
	if err == nil {
		filter := bson.M{"_id": value.ID}
		_, err := p.collection.ReplaceOne(ctx, filter, postTag)
		if err != nil {
			return err
		}
		return nil
	}
	newObj := bson.M{"post_id" : postTag.PostId, "user_id" : postTag.UserId, "hashtag" : postTag.Hashtag}
	_, err = p.collection.InsertOne(ctx, newObj)
	return err
}

func (p *postTagRepo) GetPostsByHashTag(hashTag string, ctx context.Context) ([]string, error) {
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

func (p *postTagRepo) GetPostTagsById(id string, ctx context.Context) (*domain.PostTags, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filterPostst, err := p.collection.Find(ctx, bson.M{"post_id" : id})
	if err != nil {
		return nil, err
	}

	var postTags domain.PostTags
	if err := filterPostst.Decode(&postTags); err != nil {
		return nil, err
	}

	return &postTags, err
}

func (p *postTagRepo) GetByIdAndHashTag(id string, hashTag string, ctx context.Context) (*domain.PostTags, error) {
	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var val *domain.PostTags
	err := p.collection.FindOne(ctx, bson.M{"post_id" : id, "hashtag" : hashTag}).Decode(&val)

	if err != nil {
		return val, err
	}

	return val, nil
}

func NewPostTagRepo(db *mongo.Client) PostTagRepo {
	return &postTagRepo{
		db: db,
		collection : db.Database("search_db").Collection("post_tags"),
	}
}