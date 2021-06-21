package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"search-service/dto"
	"search-service/usecase"
)

type postTagHandler struct {
	PostTagUseCase usecase.PostTagUsecase
}


type PostTagHandler interface {
	GetPostsByHashTag(ctx *gin.Context)
	SaveNewPostTag(ctx *gin.Context)
}


func (p *postTagHandler) GetPostsByHashTag(ctx *gin.Context) {

	searchedTag := ctx.Request.URL.Query().Get("searchedTag")


	postsIds, err := p.PostTagUseCase.GetPostsByHashTag(searchedTag, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "No posts with that searched hash tag")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, postsIds)
}

func (p *postTagHandler) SaveNewPostTag(ctx *gin.Context) {
	var newLocationTag dto.PostTagProfileDTO
	err := json.NewDecoder(ctx.Request.Body).Decode(&newLocationTag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	err = p.PostTagUseCase.SaveNewPostTag(newLocationTag, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"message" : "Failed to insert"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message" : "Inserted"})
}

func NewPostTagHandler(tagUsecase usecase.PostTagUsecase) PostTagHandler {
	return &postTagHandler{PostTagUseCase: tagUsecase}
}
