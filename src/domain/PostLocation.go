package domain

type PostLocation struct {
	PostId	uint64 `bson:"post_id" json:"post_id"`
	Location	Location `bson:"location" json:"location"`
}
