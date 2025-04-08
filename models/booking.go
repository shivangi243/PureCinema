package models

type Booking struct {
	Movie    string   `json:"movie"`
	Showtime string   `json:"showtime"`
	Seats    []string `json:"seats"`
}
