package model

type City struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}
