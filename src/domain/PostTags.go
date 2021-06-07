package domain

type PostTags struct {
	ID string `bson:"_id"`
	PostId	string `bson:"post_id" json:"post_id"`
	UserId string `bson:"user_id json:"user_id"`
	Hashtag	string `bson:"hashtag" json:"hashtag"`
}

func NewPostTags(postId string, userId string, hashTag string) PostTags {
	return PostTags{PostId: postId, UserId: userId, Hashtag: hashTag}
}