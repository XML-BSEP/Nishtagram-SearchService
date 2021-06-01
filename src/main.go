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

	g := gin.Default()
	g.Run("localhost:8087")



}
