package usecase

import (
	"go.mongodb.org/mongo-driver/mongo"
	"search-service/repository"
)

type PostTagUsecase interface {
	GetByPostId(id uint64) *mongo.SingleResult
	GetByTagId(id uint64) *mongo.SingleResult
}

type postTagUseCase struct {
	PostTagRepo repository.PostTagRepo
}

func (p postTagUseCase) GetByTagId(id uint64) *mongo.SingleResult {
	return p.PostTagRepo.GetByTagId(id)
}

func (p postTagUseCase) GetByPostId(id uint64) *mongo.SingleResult {
	return p.PostTagRepo.GetByPostId(id)
}

func NewPostTagUseCase(repo repository.PostTagRepo) PostTagUsecase {
	return &postTagUseCase{
		PostTagRepo: repo,
	}
}


