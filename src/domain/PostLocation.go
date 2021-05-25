package domain

type PostLocation struct {
	PostId	uint64 `json:"post_id"`
	Location	Location `json:"location"`
}
