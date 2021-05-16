package domain

type Home struct {
	ID        int64     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	UserName  string  `json:"user_name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Image     string  `json:"image"`
}