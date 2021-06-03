package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"search-service/dto"
	"search-service/usecase"
)

type locationHandler struct {
	LocationUseCase usecase.LocationUsecase
}



type LocationHandler interface {
	GetLocationsByContains(ctx *gin.Context)
	GetExactLocation(ctx *gin.Context)
}


func (l *locationHandler) GetLocationsByContains(ctx *gin.Context) {
	location := struct {
		Location string
	}{}
	decoder := json.NewDecoder(ctx.Request.Body)
	dec_err := decoder.Decode(&location)


	if dec_err != nil {
		ctx.JSON(http.StatusBadRequest, "Location decoding error")
		ctx.Abort()
		return
	}

	locations, err := l.LocationUseCase.ContainsLocation(location.Location, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "No locations with that search parameter")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": locations})

}

func (l *locationHandler) GetExactLocation(ctx *gin.Context) {
	var locationDto dto.LocationDTO
	decoder := json.NewDecoder(ctx.Request.Body)
	dec_err := decoder.Decode(&locationDto)

	if dec_err != nil {
		ctx.JSON(http.StatusBadRequest, "Location decoding error")
		ctx.Abort()
		return
	}

	location, err := l.LocationUseCase.ExactLocation(locationDto.Longitude, locationDto.Latitude, ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "No location with that search parameter")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": location})

}



func NewLocationHandler(locationUseCase usecase.LocationUsecase) LocationHandler{
	return &locationHandler{LocationUseCase: locationUseCase}
}