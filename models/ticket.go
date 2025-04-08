// models/ticket.go
package models

import "time"

type SavedTicket struct {
	Filename  string    `json:"filename"`
	Movie     string    `json:"movie"`
	Showtime  string    `json:"showtime"`
	Seats     []string  `json:"seats"`
	User      User      `json:"user"`
	Timestamp time.Time `json:"timestamp"`
}
