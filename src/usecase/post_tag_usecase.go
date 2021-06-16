package usecase

import (
	"context"
	"search-service/domain"
	"search-service/dto"
	"search-service/repository"
)

type PostTagUsecase interface {
	GetPostsByHashTag(hashTag string, ctx context.Context) (*[]dto.PostTagsDTO, error)
}

type postTagUseCase struct {
	PostTagRepo repository.PostTagRepo
}

func (p postTagUseCase) GetPostsByHashTag(hashTag string, ctx context.Context) (*[]dto.PostTagsDTO, error) {
	postsId, err := p.PostTagRepo.GetPostsByHashTag(hashTag, ctx)
	if err != nil {
		return nil, err
	}

	var postTags []domain.PostTag
	for _, postId := range postsId {
		post := p.PostTagRepo.GetPostTagById(postId, ctx)
		postTags = append(postTags, post)
	}

	var postTagsDTOs []dto.PostTagsDTO
	for _, post := range postTags {
		postTagName, err := p.PostTagRepo.GetPostsBbyHashTagName(post.Hashtag, ctx)
		if err != nil {
			return nil, err
		}

		var postProfileLocationIds []dto.PostProfileId
		for _, postLName := range *postTagName {
			postProfileLocationIds = append(postProfileLocationIds, dto.PostProfileId{PostId: postLName.PostId, ProfileId: postLName.ProfileId})
		}

		var postTag dto.PostTagsDTO
		postTag.PostProfileId = postProfileLocationIds
		postTag.Hashtag = post.Hashtag
		postTagsDTOs = AppendIfMissingTag(postTagsDTOs, postTag)
	}

	return &postTagsDTOs, nil

}

func NewPostTagUseCase(repo repository.PostTagRepo) PostTagUsecase {
	return &postTagUseCase{
		PostTagRepo: repo,
	}
}


func AppendIfMissingTag(slice []dto.PostTagsDTO, i dto.PostTagsDTO) []dto.PostTagsDTO {
	for _, ele := range slice {

		for _, elem := range ele.PostProfileId {
			for _, elem2 := range i.PostProfileId {
				if elem == elem2 {
					return slice
				}

			}
		}

	}
	return append(slice, i)
}

