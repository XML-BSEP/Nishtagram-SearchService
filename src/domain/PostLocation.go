package domain

type PostLocation struct {
	ID string `bson:"_id"`
	PostId	string `bson:"post_id" json:"post_id"`
	UserId string `bson:"user_Id json:"user_id""`
	Location	Location `bson:"location" json:"location"`
}

func NewPostLocation(postId string, userId string, location string, longitude float64, latitude float64) PostLocation {
	newLocation := Location{LocationName: location, Latitude: latitude, Longitude: longitude}
	return PostLocation{PostId: postId, UserId: userId, Location: newLocation}
}
