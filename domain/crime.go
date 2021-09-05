package domain

type Crime struct {
	ID           int64    	`json:"id"`
	LocationName string  	`json:"location_name"`
	Longitude    float64 	`json:"longitude"`
	Latitude     float64 	`json:"latitude"`
	Description  string  	`json:"description"`
	Image        string  	`json:"image"`
	Date         string  	`json:"date"`
}

type HomeCrime struct {
	Crime
	Distance 	float64		`json:"distance"`
}
