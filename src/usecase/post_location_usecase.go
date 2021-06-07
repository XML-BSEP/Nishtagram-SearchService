package usecase

import (
	"context"
	"search-service/domain"
	"search-service/repository"
)

type postLocationUsecase struct {
	 PostLocationRepo repository.PostLocationRepo
}


type PostLocationUsecase interface {
	GetPostsByExactLocation(longitude float64, latitude float64 , ctx context.Context) ([]string, error)
	GetPostsByLocationContains(location string, ctx context.Context) ([]string, error)
	Create(postLocation domain.PostLocation, ctx context.Context) error
}

func (p *postLocationUsecase) GetPostsByExactLocation(longitude float64, latitude float64, ctx context.Context) ([]string, error) {
	return p.PostLocationRepo.GetPostsByExactLocation(longitude, latitude, ctx)
}


func (p *postLocationUsecase) GetPostsByLocationContains(location string, ctx context.Context) ([]string, error) {
	return p.PostLocationRepo.GetPostsByLocationContains(location, ctx)
}

func (p *postLocationUsecase) Create(postLocation domain.PostLocation, ctx context.Context) error {
	return p.PostLocationRepo.Create(postLocation, ctx)
}

func NewPostLocationUsecase(repo repository.PostLocationRepo) PostLocationUsecase {
	return &postLocationUsecase{ PostLocationRepo: repo}
}