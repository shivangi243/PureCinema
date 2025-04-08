package db

import (
	"cinema/models"
	"encoding/json"
	"fmt"
	"sync"
)

var seatMutex sync.Mutex // Global mutex for thread-safe booking

// SaveBooking stores the booking details into the database safely
func SaveBooking(b models.Booking) error {
	seatMutex.Lock()
	defer seatMutex.Unlock()

	// Step 1: Check if any of the seats are already booked
	bookedList, err := GetBookedSeats(b.Movie, b.Showtime)
	if err != nil {
		return fmt.Errorf("failed to check seat availability: %v", err)
	}

	// Convert to map for quick lookup
	bookedSeats := make(map[string]bool)
	for _, s := range bookedList {
		bookedSeats[s] = true
	}

	for _, seat := range b.Seats {
		if bookedSeats[seat] {
			return fmt.Errorf("seat %s is already booked", seat)
		}
	}

	// Step 2: Marshal and insert booking
	seatsJSON, err := json.Marshal(b.Seats)
	if err != nil {
		return fmt.Errorf("failed to marshal seats: %v", err)
	}

	_, err = Conn.Exec("INSERT INTO bookings (movie, showtime, seats) VALUES (?, ?, ?)", b.Movie, b.Showtime, seatsJSON)
	if err != nil {
		return fmt.Errorf("failed to save booking: %v", err)
	}

	return nil
}
