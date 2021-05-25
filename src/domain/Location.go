package domain

type Location struct {
	LocationId	uint64 `json:"location_id"`
	Longitude	float64 `json:"longitude"`
	Latitude	float64 `json:"latitude"`
}
