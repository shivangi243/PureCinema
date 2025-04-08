package tests

import (
	"cinema/db"
	"cinema/models"
	"encoding/json"
	"testing"
)

func TestUnmarshalSeats(t *testing.T) {
	jsonData := `["A1", "A2", "B3"]`

	var seats []string
	err := json.Unmarshal([]byte(jsonData), &seats)
	if err != nil {
		t.Errorf("Unmarshal failed: %v", err)
	}

	if len(seats) != 3 {
		t.Errorf("Expected 3 seats, got %d", len(seats))
	}
}

func TestSaveBooking(t *testing.T) {
	// Connect to the MySQL database
	db.Connect()

	booking := models.Booking{
		Movie:    "Unit Test Movie",
		Showtime: "11:11 AM",
		Seats:    []string{"Z1", "Z2"},
	}

	err := db.SaveBooking(booking)
	if err != nil {
		t.Errorf("SaveBooking failed: %v", err)
	}
}
