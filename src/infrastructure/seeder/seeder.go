package seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"search-service/domain"
)

func DropDatabase(db string, mongoCli *mongo.Client, ctx context.Context){
	err := mongoCli.Database(db).Drop(ctx)
	if err != nil {
		return
	}
}

func SeedData(db string, mongoCli *mongo.Client, ctx context.Context) {
	DropDatabase(db, mongoCli, ctx)

	if cnt,_ := mongoCli.Database(db).Collection("locations").EstimatedDocumentCount(ctx, nil); cnt == 0 {
		locationCollection := mongoCli.Database(db).Collection("locations")
		seedLocation(locationCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("post_locations").EstimatedDocumentCount(ctx, nil); cnt == 0 {
		postLocationCollection := mongoCli.Database(db).Collection("post_locations")
		seedPostLocations(postLocationCollection, ctx)
	}


	if cnt,_ := mongoCli.Database(db).Collection("post_tags").EstimatedDocumentCount(ctx, nil); cnt == 0 {
		postTags := mongoCli.Database(db).Collection("post_tags")
		seedPostTags(postTags, ctx)
	}

}

func seedPostTags(tags *mongo.Collection, ctx context.Context) {
	_, err := tags.InsertMany(ctx, []interface{} {
		bson.D{
			{"post_id", "123451"},
			{"hashtag", "bff"},
		},
		bson.D{
			{"post_id", "123452"},
			{"hashtag", "tbt"},
		},
		bson.D{
			{"post_id", "123453"},
			{"hashtag", "idegasnamax"},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}


func seedLocation(locationCollection *mongo.Collection, ctx context.Context) {

	_, err := locationCollection.InsertMany(ctx, []interface{} {
		bson.D{
			{"location", "Cara Dušana 5, Novi Sad, Srbija"},
			{"longitude", 50},
			{"latitude", 60},
		},
		bson.D{
			{"location", "Hadži Ruvimova 10, Novi Sad, Srbija"},
			{"longitude", 500},
			{"latitude", 600},
		},
		bson.D{
			{"location", " Gospodara Vučića BB, Beograd, Srbija"},
			{"longitude", 1},
			{"latitude", 2},
		},

	})

	if err != nil {
		log.Fatal(err)
	}
}


func seedPostLocations(postLocationCollection *mongo.Collection, ctx context.Context) {
	location1 := domain.Location{Location: "Cara Dušana 5, Novi Sad, Srbija", Longitude: 50, Latitude: 60}
	location2 := domain.Location{Location: "Hadži Ruvimova 10, Novi Sad, Srbija", Longitude: 500, Latitude: 600}
	location3 := domain.Location{Location: " Gospodara Vučića BB, Beograd, Srbija", Longitude: 1, Latitude: 2}

	_, err := postLocationCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"post_id", "123451"},
			{"location", location1},

		},
		bson.D{
			{"post_id", "123452"},
			{"location", location2},

		},
		bson.D{
			{"post_id", "123453"},
			{"location", location3},

		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
