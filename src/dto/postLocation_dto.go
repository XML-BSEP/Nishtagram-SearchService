package dto

import "search-service/domain"

type PostLocationDTO struct {
	PostId	string `bson:"post_id" json:"post_id"`
	Location string `bson:"location" json:"location"`
}

type PostLocationsDTO struct {
	PostId	[]string `bson:"post_id" json:"post_id"`
	Location string `bson:"location" json:"location"`
}

func NewPostLocationDTO(postLocation domain.PostLocation) PostLocationDTO {
	return PostLocationDTO{PostId: postLocation.PostId, Location: postLocation.Location.Location}
}
