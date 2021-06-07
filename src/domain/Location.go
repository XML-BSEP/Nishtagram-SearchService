package domain

type Location struct {
	LocationName	string `bson:"location" json:"location"`
	Longitude	float64 `bson:"longitude" json:"longitude"`
	Latitude	float64 `bson:"latitude" json:"latitude"`
}
