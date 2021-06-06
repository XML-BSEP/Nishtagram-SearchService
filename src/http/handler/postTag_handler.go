package handler

import (
	"encoding/json"
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
	hashTag := struct {
		HashTag string `json:"hash_tag"`
	}{}

	decoder := json.NewDecoder(ctx.Request.Body)
	dec_err := decoder.Decode(&hashTag)

	if dec_err != nil {
		ctx.JSON(http.StatusBadRequest, "Post tag decoding error")
		ctx.Abort()
		return
	}

	postsIds, err := p.PostTagUseCase.GetPostsByHashTag(hashTag.HashTag, ctx)
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
