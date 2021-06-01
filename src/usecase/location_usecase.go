package usecase

import (
	"context"
	"search-service/domain"
	"search-service/repository"
)

type locationUsecase struct {
	LocationRepo repository.LocationRepo
}


type LocationUsecase interface {
	ContainsLocation(location string, ctx context.Context) ([]domain.Location, error)
	ExactLocation(longitude float64, latitude float64, ctx context.Context) (domain.Location, error)
}


func (l locationUsecase) ExactLocation(longitude float64, latitude float64, ctx context.Context) (domain.Location, error) {
	return  l.LocationRepo.ExactLocation(longitude, latitude, ctx)
}

func (l locationUsecase) ContainsLocation(location string, ctx context.Context) ([]domain.Location, error) {
	return  l.LocationRepo.ContainsLocation(location, ctx)
}


func NewLocationUsecase(repo repository.LocationRepo) LocationUsecase {
	return &locationUsecase{ LocationRepo: repo}
}
