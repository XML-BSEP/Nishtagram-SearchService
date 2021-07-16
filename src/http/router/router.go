package router

import (
	"github.com/gin-gonic/gin"
	"search-service/http/middleware/prometheus_middleware"
	"search-service/interactor"
)

func NewRouter(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()
	requestCoutnter := prometheus_middleware.GetHttpRequestsCounter()
	router.Use(prometheus_middleware.PrometheusMiddleware(requestCoutnter))

	router.GET("/getLocationsByContaining", handler.GetLocationsByContains)
	router.GET("/getExactLocation", handler.GetExactLocation)
	router.GET("/getPostLocationsByLocationContaining", handler.GetPostsByLocationContains)
	router.GET("/getPostsByExactLocation", handler.GetPostsByExactLocation)
	router.GET("/getPostsByTag", handler.GetPostsByHashTag)
	router.POST("/saveNewPostLocation", handler.SaveNewPostLocation)
	router.POST("/saveNewPostTag", handler.SaveNewPostTag)
	router.GET("/metrics", prometheus_middleware.PrometheusGinHandler())

	return router
}
