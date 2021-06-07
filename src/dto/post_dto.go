package dto

type PostDto struct {
	PostId string `json:"post_id"`
	UserId string `json:"user_id"`
	HashTags []string `json:"hash_tags"`
	LocationName string `json:"location"`
}
