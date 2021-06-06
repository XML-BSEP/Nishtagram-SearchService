package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"search-service/dto"
	"search-service/usecase"
)

type postLocationHandler struct {
	PostLocationUseCase usecase.PostLocationUsecase
}


type PostLocationHandler interface {
	GetPostsByExactLocation(ctx *gin.Context)
	GetPostsByLocationContains(ctx *gin.Context)

}

func (p *postLocationHandler) GetPostsByLocationContains(ctx *gin.Context) {
	location := struct {
		Location string
	}{}

	decoder := json.NewDecoder(ctx.Request.Body)
	dec_err := decoder.Decode(&location)

	if dec_err != nil {
		ctx.JSON(http.StatusBadRequest, "Post location decoding error")
		ctx.Abort()
		return
	}

	postsIds, err := p.PostLocationUseCase.GetPostsByLocationContains(location.Location, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "No posts with that search parameter")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": postsIds})
}

func (p *postLocationHandler) GetPostsByExactLocation(ctx *gin.Context) {
	var postLocationDto dto.PostLocationDTO
	decoder := json.NewDecoder(ctx.Request.Body)
	dec_err := decoder.Decode(&postLocationDto)

	if dec_err != nil {
		ctx.JSON(http.StatusBadRequest, "Post location decoding error")
		ctx.Abort()
		return
	}

	postLocations, err := p.PostLocationUseCase.GetPostsByExactLocation(postLocationDto.Location.Latitude, postLocationDto.Location.Longitude, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "No post locations with that longitude-latitude")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data" : postLocations})
}

func NewPostLocationHandler(locationUsecase usecase.PostLocationUsecase) PostLocationHandler {
	return &postLocationHandler{PostLocationUseCase: locationUsecase}
}
