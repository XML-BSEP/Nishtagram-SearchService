package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"search-service/usecase"
)

type postTagHandler struct {
	PostTagUseCase usecase.PostTagUsecase
}


type PostTagHandler interface {
	GetPostsByHashTag(ctx *gin.Context)
}


func (p *postTagHandler) GetPostsByHashTag(ctx *gin.Context) {


	hashTag := ctx.Request.URL.Query().Get("hashTag")

	postsIds, err := p.PostTagUseCase.GetPostsByHashTag(hashTag, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "No posts with that searched hash tag")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": postsIds})
}

func NewPostTagHandler(tagUsecase usecase.PostTagUsecase) PostTagHandler {
	return &postTagHandler{PostTagUseCase: tagUsecase}
}
