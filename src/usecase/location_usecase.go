package usecase

import (
	"go.mongodb.org/mongo-driver/mongo"
	"search-service/repository"
)

type locationUsecase struct {
	LocationRepo repository.LocationRepo
}

type LocationUsecase interface {
	GetById(id uint64) *mongo.SingleResult
}

func (l locationUsecase) GetById(id uint64) *mongo.SingleResult {
	return l.LocationRepo.GetById(id)
}

func NewLocationUsecase(repo repository.LocationRepo) LocationUsecase {
	return &locationUsecase{ LocationRepo: repo}
}
