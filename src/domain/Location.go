package domain

type Location struct {
	LocationId	uint64 `bson:"location_id" json:"location_id"`
	Longitude	float64 `bson:"longitude" json:"longitude"`
	Latitude	float64 `bson:"latitude" json:"latitude"`
}
