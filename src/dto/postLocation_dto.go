package dto

import "search-service/domain"

type PostLocationDTO struct {
	PostId	string `bson:"post_id" json:"post_id"`
	Location string `bson:"location" json:"location"`
}
/*
type PostLocationsDTO struct {
	PostId	[]string `bson:"post_id" json:"post_id"`
	ProfileId	[]string `bson:"profile_id" json:"profile_id"`
	Location string `bson:"location" json:"location"`
}*/

type PostProfileId struct {
	PostId string `bson:"post_id" json:"post_id"`
	ProfileId string `bson:"profile_id" json:"profile_id"`
}

type PostLocationsDTO struct {
	PostProfileId	[]PostProfileId `bson:"post_profile_id" json:"post_profile_id"`
	Location string `bson:"location" json:"location"`
}

func NewPostLocationDTO(postLocation domain.PostLocation) PostLocationDTO {
	return PostLocationDTO{PostId: postLocation.PostId, Location: postLocation.Location.Location}
}
