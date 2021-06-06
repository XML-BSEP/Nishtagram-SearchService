package dto

type LocationDTO struct {
	Location	string `bson:"location" json:"location"`
	Longitude	float64 `bson:"longitude" json:"longitude"`
	Latitude	float64 `bson:"latitude" json:"latitude"`
}

