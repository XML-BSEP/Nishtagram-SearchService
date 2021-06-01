package main

import (
	"github.com/gin-gonic/gin"
	"search-service/infrastructure/mongo"
	"search-service/infrastructure/seeder"
)

func main() {

	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	println(db)
	seeder.SeedData(db, mongoCli, ctx)


	//location test
	/*
	locationRepo := repository.NewLocationRepo(mongoCli)
	locationService := usecase.NewLocationUsecase(locationRepo)
	res := locationService.GetById(111111)
	var lokacije domain.Location
	decodeErr := res.Decode(&lokacije)
	if decodeErr != nil {
		println("ovde sam")
	}
	println(lokacije.LocationId)*/

	//postLocation test
	/*
	postLocationRepo := repository.NewPostLocationRepo(mongoCli)
	postLocationService := usecase.NewPostLocationUsecase(postLocationRepo)
	res := postLocationService.GetById(123451)
	var postLoc domain.PostLocation
	decodeErr := res.Decode(&postLoc)
	if decodeErr != nil {
		println("ovde sam")
	}
	println(postLoc.PostId, postLoc.Location.LocationId)*/

	//postTag test
	/*
	postTagRepo := repository.NewPostTagRepo(mongoCli)
	postTagService := usecase.NewPostTagUseCase(postTagRepo)

	//resPost := postTagService.GetByPostId(1231)
	resTag := postTagService.GetByTagId(12322)
	var postTag domain.PostTags

	decodeErrPost := resTag.Decode(&postTag)
	if decodeErrPost != nil {
		println("ovde sam")
	}
	println(postTag.PostId)*/


	g := gin.Default()
	g.Run("localhost:8087")



}
