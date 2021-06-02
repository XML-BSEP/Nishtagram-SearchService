package dto

type PostLocationDTO struct {
	PostId	string `bson:"post_id" json:"post_id"`
	Location	LocationDTO `bson:"location" json:"location"`
}
