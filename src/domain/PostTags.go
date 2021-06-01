package domain

type PostTags struct {
	PostId	uint64 `bson:"post_id" json:"post_id"`
	HashtagId	uint64 `bson:"hashtag_id" json:"hashtag_id"`
}
