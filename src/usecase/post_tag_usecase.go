package usecase

import (
	"search-service/repository"
)

type PostTagUsecase interface {
}

type postTagUseCase struct {
	PostTagRepo repository.PostTagRepo
}



func NewPostTagUseCase(repo repository.PostTagRepo) PostTagUsecase {
	return &postTagUseCase{
		PostTagRepo: repo,
	}
}


