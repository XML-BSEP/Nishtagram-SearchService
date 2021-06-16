package dto

type PostLocationExactDTO struct {
	PostId	string `bson:"post_id" json:"post_id"`
	Latitude float64 `bson:"latitude" json:"latitude"`
	Longitude float64 `bson:"longitude" json:"longitude"`
}
