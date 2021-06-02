package router

import (
	"github.com/gin-gonic/gin"
	"search-service/interactor"
)

func NewRouter(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/getLocationsByContaining", handler.GetLocationsByContains)
	router.GET("/getExactLocation", handler.GetExactLocation)

	return router
}
