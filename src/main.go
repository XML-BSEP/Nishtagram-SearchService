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

	res, err := locationService.ContainsLocation("srb", ctx)
	if err != nil {
		fmt.Println(err)
	}
	for _, location := range res {
		fmt.Println(location)
	}

	res, err := locationService.ExactLocation(1, 2, ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res) */

	//postLocation test
	/*
	postLocationRepo := repository.NewPostLocationRepo(mongoCli)
	postLocationService := usecase.NewPostLocationUsecase(postLocationRepo)
	//res, err := postLocationService.GetPostsByExactLocation(1, 2, ctx)

	//if err != nil {
//		fmt.Println(err)
//	}

	res, err := postLocationService.GetPostsByLocationContains("beograd", ctx)
	if err != nil {
		fmt.Println(err)
	}
	for _, post := range res {
		fmt.Println(post)
	}*/



	//postTag test
	/*
	postTagRepo := repository.NewPostTagRepo(mongoCli)
	postTagService := usecase.NewPostTagUseCase(postTagRepo)

	//resPost := postTagService.GetByPostId(1231)
	resTag, err := postTagService.GetPostsByHashTag("tbt", ctx)
	if err != nil {
		fmt.Println(err)
	}
	for _, post := range resTag {
		fmt.Println(post)
	}*/


	g := gin.Default()
	g.Run("localhost:8087")



}
