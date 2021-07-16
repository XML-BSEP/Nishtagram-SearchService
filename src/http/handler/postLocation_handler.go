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
	SaveNewPostLocation(ctx *gin.Context)

}

func (p *postLocationHandler) GetPostsByLocationContains(ctx *gin.Context) {

	searchedLocation := ctx.Request.URL.Query().Get("searchedLocation")

/*
	postsIds, err := p.PostLocationUseCase.GetPostsByLocationContains(searchedLocation, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "No posts with that search parameter")
		ctx.Abort()
		return
	}
*/
	posts, err := p.PostLocationUseCase.GetPostsAndLocationByLocationContaining(searchedLocation, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "No posts with that search parameter")
		ctx.Abort()
		return
	}


	ctx.JSON(http.StatusOK, posts)
}

func (p *postLocationHandler) GetPostsByExactLocation(ctx *gin.Context) {
	var postLocationDto dto.PostLocationExactDTO
	decoder := json.NewDecoder(ctx.Request.Body)
	dec_err := decoder.Decode(&postLocationDto)

	if dec_err != nil {
		ctx.JSON(http.StatusBadRequest, "Post location decoding error")
		ctx.Abort()
		return
	}


	postLocations, err := p.PostLocationUseCase.GetPostsByExactLocation(postLocationDto.Latitude, postLocationDto.Longitude, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "No post locations with that longitude-latitude")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data" : postLocations})
}

func (p *postLocationHandler) SaveNewPostLocation(ctx *gin.Context) {
	var newLocationPost dto.PostLocationProfileDTO
	err := json.NewDecoder(ctx.Request.Body).Decode(&newLocationPost)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	err = p.PostLocationUseCase.SaveNewPostLocation(newLocationPost, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"message" : "Failed to insert"})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message" : "Inserted"})
}

func NewPostLocationHandler(locationUsecase usecase.PostLocationUsecase) PostLocationHandler {
	return &postLocationHandler{PostLocationUseCase: locationUsecase}
}
