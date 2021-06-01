package usecase

import (
	"search-service/repository"
)

type postLocationUsecase struct {
	 PostLocationRepo repository.PostLocationRepo
}

type PostLocationUsecase interface {

}



func NewPostLocationUsecase(repo repository.PostLocationRepo) PostLocationUsecase {
	return &postLocationUsecase{ PostLocationRepo: repo}
}