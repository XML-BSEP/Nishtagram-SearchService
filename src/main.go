package main

import (
	router2 "search-service/http/router"
	"search-service/infrastructure/mongo"
	"search-service/infrastructure/seeder"
	interactor2 "search-service/interactor"
)

func main() {

	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	println(db)
	seeder.SeedData(db, mongoCli, ctx)


	interactor := interactor2.NewInteractor(mongoCli)
	appHandler := interactor.NewAppHandler()

	router := router2.NewRouter(appHandler)
	router.Run(":8087")



	//g := gin.Default()
	//g.Run("localhost:8087")



}
