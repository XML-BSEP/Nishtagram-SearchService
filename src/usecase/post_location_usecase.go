package usecase

import (
	"go.mongodb.org/mongo-driver/mongo"
	"search-service/repository"
)

type postLocationUsecase struct {
	 PostLocationRepo repository.PostLocationRepo
}

type PostLocationUsecase interface {
	GetById(id uint64) *mongo.SingleResult
}


func (p postLocationUsecase) GetById(id uint64) *mongo.SingleResult {
	return  p.PostLocationRepo.GetById(id)
}


func NewPostLocationUsecase(repo repository.PostLocationRepo) PostLocationUsecase {
	return &postLocationUsecase{ PostLocationRepo: repo}
}