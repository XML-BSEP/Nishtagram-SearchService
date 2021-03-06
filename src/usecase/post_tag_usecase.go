package usecase

import (
	"context"
	"search-service/domain"
	"search-service/dto"
	"search-service/repository"
)

type PostTagUsecase interface {
	GetPostsByHashTag(hashTag string, ctx context.Context) (*[]dto.PostTagsDTO, error)
	SaveNewPostTag(location dto.PostTagProfileDTO, ctx context.Context) error
}

type postTagUseCase struct {
	PostTagRepo repository.PostTagRepo
}

func (p postTagUseCase) GetPostsByHashTag(hashTag string, ctx context.Context) (*[]dto.PostTagsDTO, error) {
	postsId, err := p.PostTagRepo.GetPostsByHashTag(hashTag, ctx)
	if err != nil {
		return nil, err
	}


	var postTagsDTOs []dto.PostTagsDTO
	for _, post := range postsId {
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
		postTagsDTOs = append(postTagsDTOs, postTag)
		//postTagsDTOs = AppendIfMissingTag(postTagsDTOs, postTag)
	}

	var retVal []dto.PostTagsDTO
	var oneTag dto.PostTagsDTO
	for _, p := range postTagsDTOs {
		toAdd := true
		for _, o := range retVal {
			if o.Hashtag == p.Hashtag {
				toAdd = false
				break
			}
		}
		if toAdd {
			oneTag.Hashtag = p.Hashtag
			oneTag.PostProfileId = p.PostProfileId
			retVal = append(retVal, oneTag)
		}
	}

	return &retVal, nil

}

func (p postTagUseCase) SaveNewPostTag(tag dto.PostTagProfileDTO, ctx context.Context) error {


	for _, s := range tag.Hashtag {
		var postTag domain.PostTag
		postTag.PostId = tag.PostId
		postTag.Hashtag = s
		postTag.ProfileId = tag.ProfileId

		error := p.PostTagRepo.SaveNewPostTag(postTag, ctx)
		if error != nil {
			continue
		}
	}

	return nil
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

