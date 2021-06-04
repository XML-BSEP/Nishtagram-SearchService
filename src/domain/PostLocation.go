package domain

type PostLocation struct {
	PostId	string `bson:"post_id" json:"post_id"`
	Location	Location `bson:"location" json:"location"`
}
