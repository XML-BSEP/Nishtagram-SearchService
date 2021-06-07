package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"search-service/domain"
	"search-service/dto"
	"search-service/usecase"
)

type postHandler struct {
	PostTagUsecase usecase.PostTagUsecase
	PostLocationUsecase usecase.PostLocationUsecase
}

type PostHandler interface {
	InsertPost(ctx *gin.Context)
}

func NewPostHandler(postTagUsecase usecase.PostTagUsecase, postLocationUsecase usecase.PostLocationUsecase) PostHandler{
	return &postHandler{postTagUsecase, postLocationUsecase}
}

func (p *postHandler) InsertPost(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)

	var postDto dto.PostDto

	if err := decoder.Decode(&postDto); err != nil {
		ctx.JSON(500, gin.H{"message" : "Error decoding struct"})
		return
	}

	for _, hashTag := range postDto.HashTags {
		postTag := domain.NewPostTags(postDto.PostId, postDto.UserId, hashTag)
		if err := p.PostTagUsecase.Create(postTag, ctx); err != nil {
			ctx.JSON(400, gin.H{"message" : "Can not save post tags"})
			return
		}
	}

	if err := p.PostLocationUsecase.Create(domain.NewPostLocation(postDto.PostId, postDto.UserId, postDto.LocationName, 0.0, 0.0), ctx); err != nil {
		ctx.JSON(400, gin.H{"message" : "Can not save post location"})
		return
	}

	ctx.JSON(200, gin.H{"message" : "Post saved successful"})

}