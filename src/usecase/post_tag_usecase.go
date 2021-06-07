package usecase

import (
	"context"
	"search-service/domain"
	"search-service/repository"
)

type PostTagUsecase interface {
	GetPostsByHashTag(hashTag string, ctx context.Context) ([]string, error)
	Create(postTag domain.PostTags, ctx context.Context) error
}

type postTagUseCase struct {
	PostTagRepo repository.PostTagRepo
}

func (p *postTagUseCase) Create(postTag domain.PostTags, ctx context.Context) error {
	return p.PostTagRepo.Create(postTag, ctx)
}

func (p *postTagUseCase) GetPostsByHashTag(hashTag string, ctx context.Context) ([]string, error) {
	return  p.PostTagRepo.GetPostsByHashTag(hashTag, ctx)
}

func NewPostTagUseCase(repo repository.PostTagRepo) PostTagUsecase {
	return &postTagUseCase{
		PostTagRepo: repo,
	}
}


