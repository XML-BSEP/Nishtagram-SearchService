package seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"search-service/domain"
)

func DropDatabase(db string, mongoCli *mongo.Client, ctx *context.Context){
	err := mongoCli.Database(db).Drop(*ctx)
	if err != nil {
		return
	}
}

func SeedData(db string, mongoCli *mongo.Client, ctx *context.Context) {
	DropDatabase(db, mongoCli, ctx)

	if cnt,_ := mongoCli.Database(db).Collection("locations").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		locationCollection := mongoCli.Database(db).Collection("locations")
		seedLocation(locationCollection, ctx)
	}

	if cnt,_ := mongoCli.Database(db).Collection("post_locations").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		postLocationCollection := mongoCli.Database(db).Collection("post_locations")
		seedPostLocations(postLocationCollection, ctx)
	}


	if cnt,_ := mongoCli.Database(db).Collection("post_tags").EstimatedDocumentCount(*ctx, nil); cnt == 0 {
		postTags := mongoCli.Database(db).Collection("post_tags")
		seedPostTags(postTags, ctx)
	}

}

func seedPostTags(tags *mongo.Collection, ctx *context.Context) {
	_, err := tags.InsertMany(*ctx, []interface{} {
		bson.D{
			{"post_id", "1231"},
			{"hashtag_id", "1231"},
		},
		bson.D{
			{"post_id", "1232"},
			{"hashtag_id", "1232"},
		},
		bson.D{
			{"post_id", "1233"},
			{"hashtag_id", "1233"},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}


func seedLocation(locationCollection *mongo.Collection, ctx *context.Context) {

	_, err := locationCollection.InsertMany(*ctx, []interface{} {
		bson.D{
			{"location_id", "111111"},
			{"longitude", "50"},
			{"latitude", "60"},
		},
		bson.D{
			{"location_id", "111112"},
			{"longitude", "500"},
			{"latitude", "600"},
		},
		bson.D{
			{"location_id", "111113"},
			{"longitude", "450"},
			{"latitude", "460"},
		},
		bson.D{
			{"location_id", "111114"},
			{"longitude", "1"},
			{"latitude", "2"},
		},
		bson.D{
			{"location_id", "111115"},
			{"longitude", "50.10"},
			{"latitude", "60.131212"},
		},

	})

	if err != nil {
		log.Fatal(err)
	}
}


func seedPostLocations(postLocationCollection *mongo.Collection, ctx *context.Context) {
	location1 := domain.Location{LocationId: 111111, Longitude: 50, Latitude: 60}
	location2 := domain.Location{LocationId: 111112, Longitude: 500, Latitude: 600}
	location3 := domain.Location{LocationId: 111113, Longitude: 450, Latitude: 460}
	location4 := domain.Location{LocationId: 111114, Longitude: 1, Latitude: 2}
	location5 := domain.Location{LocationId: 111115, Longitude: 50.10, Latitude: 60.131212}

	_, err := postLocationCollection.InsertMany(*ctx, []interface{}{
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
		bson.D{
			{"post_id", "123454"},
			{"location", location4},

		},
		bson.D{
			{"post_id", "123455"},
			{"location", location5},

		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
