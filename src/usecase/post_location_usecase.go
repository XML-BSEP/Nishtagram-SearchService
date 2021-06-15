package usecase

import (
	"context"
	"search-service/domain"
	"search-service/dto"
	"search-service/repository"
)

type postLocationUsecase struct {
	 PostLocationRepo repository.PostLocationRepo
}


type PostLocationUsecase interface {
	GetPostsByExactLocation(longitude float64, latitude float64 , ctx context.Context) ([]string, error)
	GetPostsByLocationContains(location string, ctx context.Context) ([]string, error)
	GetPostsAndLocationByLocationContaining(location string, ctx context.Context) (*[]dto.PostLocationsDTO, error)
}

func (p postLocationUsecase) GetPostsByExactLocation(longitude float64, latitude float64, ctx context.Context) ([]string, error) {
	return p.PostLocationRepo.GetPostsByExactLocation(longitude, latitude, ctx)
}


func (p postLocationUsecase) GetPostsByLocationContains(location string, ctx context.Context) ([]string, error) {
	return p.PostLocationRepo.GetPostsByLocationContains(location, ctx)
}


func (p postLocationUsecase) GetPostsAndLocationByLocationContaining(location string, ctx context.Context) (*[]dto.PostLocationsDTO, error) {
	postsId, err := p.GetPostsByLocationContains(location, ctx)
	if err != nil {
		return nil, err
	}

	var postLocations []domain.PostLocation
	for _, postId := range postsId {
		post  := p.PostLocationRepo.GetPostLocationById(postId, ctx)
		postLocations = append(postLocations, post)
	}


	var postLocationsDTOs []dto.PostLocationsDTO
	for _, post := range postLocations {
		postLocationByName, err := p.PostLocationRepo.GetPostsByLocationName(post.Location.Location, ctx)
		if err != nil {
			return nil, err
		}

		var postLocationIds []string
		for _, postLName := range *postLocationByName { //jedan iz beograd cara lazara
			postLocationIds = append(postLocationIds, postLName.PostId)
		}

		var postLocation dto.PostLocationsDTO
		postLocation.PostId = postLocationIds
		postLocation.Location = post.Location.Location
		postLocationsDTOs = AppendIfMissing(postLocationsDTOs, postLocation)
	}

	return &postLocationsDTOs, nil

}

func NewPostLocationUsecase(repo repository.PostLocationRepo) PostLocationUsecase {
	return &postLocationUsecase{ PostLocationRepo: repo}
}

func AppendIfMissing(slice []dto.PostLocationsDTO, i dto.PostLocationsDTO) []dto.PostLocationsDTO {
	for _, ele := range slice {

		for _, elem := range ele.PostId {
			for _, elem2 := range i.PostId {
				if elem == elem2 {
					return slice
				}

			}
		}

	}
	return append(slice, i)
}