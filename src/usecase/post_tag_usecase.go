package usecase

import (
	"context"
	"search-service/repository"
)

type PostTagUsecase interface {
	GetPostsByHashTag(hashTag string, ctx context.Context) ([]string, error)
}

type postTagUseCase struct {
	PostTagRepo repository.PostTagRepo
}

func (p postTagUseCase) GetPostsByHashTag(hashTag string, ctx context.Context) ([]string, error) {
	return  p.PostTagRepo.GetPostsByHashTag(hashTag, ctx)
}

func NewPostTagUseCase(repo repository.PostTagRepo) PostTagUsecase {
	return &postTagUseCase{
		PostTagRepo: repo,
	}
}


