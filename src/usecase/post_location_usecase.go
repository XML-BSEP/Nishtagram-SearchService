package usecase

import (
	"context"
	"search-service/repository"
)

type postLocationUsecase struct {
	 PostLocationRepo repository.PostLocationRepo
}


type PostLocationUsecase interface {
	GetPostsByExactLocation(longitude float64, latitude float64 , ctx context.Context) ([]string, error)
	GetPostsByLocationContains(location string, ctx context.Context) ([]string, error)
}

func (p postLocationUsecase) GetPostsByExactLocation(longitude float64, latitude float64, ctx context.Context) ([]string, error) {
	return p.PostLocationRepo.GetPostsByExactLocation(longitude, latitude, ctx)
}


func (p postLocationUsecase) GetPostsByLocationContains(location string, ctx context.Context) ([]string, error) {
	return p.PostLocationRepo.GetPostsByLocationContains(location, ctx)
}


func NewPostLocationUsecase(repo repository.PostLocationRepo) PostLocationUsecase {
	return &postLocationUsecase{ PostLocationRepo: repo}
}