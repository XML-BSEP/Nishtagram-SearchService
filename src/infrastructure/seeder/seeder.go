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
			{"post_id", "4752f49f-3011-44af-9c62-2a6f4086233d"},
			{"profile_id", "e2b5f92e-c31b-11eb-8529-0242ac130003"},
			{"hashtag", "tbt"},
		},
		bson.D{
			{"post_id", "d459e0f2-ab61-48e8-a593-29933ce99525"},
			{"profile_id", "424935b1-766c-4f99-b306-9263731518bc"},
			{"hashtag", "idegasnamax"},
		},
		bson.D{
			{"post_id", "4752f49f-3011-44af-9c62-2a6f4086233d"},
			{"profile_id", "e2b5f92e-c31b-11eb-8529-0242ac130003"},
			{"hashtag", "idegasnamax"},
		},
		bson.D{
			{"post_id", "1ea5b7bc-94eb-40c0-98fd-7858e197e3b2"},
			{"profile_id", "a2c2f993-dc32-4a82-82ed-a5f6866f7d03"},
			{"hashtag", "idegasnamax"},
		},
		bson.D{
			{"post_id", "adfee6f4-fe45-40ad-8f8e-760ec861a35e"},
			{"profile_id", "43420055-3174-4c2a-9823-a8f060d644c3"},
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
			{"location", "KI"},
			{"longitude", 50},
			{"latitude", 60},
		},
		bson.D{
			{"location", "NS"},
			{"longitude", 500},
			{"latitude", 600},
		},
		bson.D{
			{"location", " SM"},
			{"longitude", 1},
			{"latitude", 2},
		},

	})

	if err != nil {
		log.Fatal(err)
	}
}


func seedPostLocations(postLocationCollection *mongo.Collection, ctx context.Context) {
	location1 := domain.Location{Location: "KI", Longitude: 50, Latitude: 60}
	location2 := domain.Location{Location: "NS", Longitude: 500, Latitude: 600}
	location3 := domain.Location{Location: "SM", Longitude: 1, Latitude: 2}

	_, err := postLocationCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"post_id", "4752f49f-3011-44af-9c62-2a6f4086233d"},
			{"profile_id", "e2b5f92e-c31b-11eb-8529-0242ac130003"},
			{"location", location1},

		},
		bson.D{
			{"post_id", "d459e0f2-ab61-48e8-a593-29933ce99525"},
			{"profile_id", "424935b1-766c-4f99-b306-9263731518bc"},
			{"location", location1},

		},
		bson.D{
			{"post_id", "1ea5b7bc-94eb-40c0-98fd-7858e197e3b2"},
			{"profile_id", "a2c2f993-dc32-4a82-82ed-a5f6866f7d03"},
			{"location", location3},

		},
		bson.D{
			{"post_id", "adfee6f4-fe45-40ad-8f8e-760ec861a35e"},
			{"profile_id", "43420055-3174-4c2a-9823-a8f060d644c3"},
			{"location", location2},

		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
