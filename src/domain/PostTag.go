package domain

type PostTag struct {
	PostId	string `bson:"post_id" json:"post_id"`
	ProfileId	string `bson:"profile_id" json:"profile_id"`
	Hashtag	string `bson:"hashtag" json:"hashtag"`
}
