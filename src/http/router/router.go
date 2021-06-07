package router

import (
	"github.com/gin-gonic/gin"
	"search-service/interactor"
)

func NewRouter(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/getLocationsByContaining", handler.GetLocationsByContains)
	router.GET("/getExactLocation", handler.GetExactLocation)
	router.GET("/getPostsByLocationContaining", handler.GetPostsByLocationContains)
	router.GET("/getPostsByExactLocation", handler.GetPostsByExactLocation)
	router.GET("/getPostsByHashTag", handler.GetPostsByHashTag)
	router.POST("/insertPost", handler.InsertPost)

	return router
}
