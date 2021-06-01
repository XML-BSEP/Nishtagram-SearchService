package domain

type PostTags struct {
	PostId	string `bson:"post_id" json:"post_id"`
	Hashtag	string `bson:"hashtag" json:"hashtag"`
}
