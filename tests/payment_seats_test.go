package tests

import (
	"cinema/db"
	"cinema/models"
	"encoding/json"
	"errors"
	"sync"
	"testing"
	"time"
)

// Mutex to prevent race conditions in concurrent tests
var bookingMu sync.Mutex

// ✅ Test saving and retrieving a booking from DB
func TestSaveBookingAndRetrieve(t *testing.T) {
	bookingMu.Lock()
	defer bookingMu.Unlock()

	booking := models.Booking{
		Movie:    "Dune",
		Showtime: "10:00 PM",
		Seats:    []string{"A1", "A2"},
	}

	err := db.SaveBooking(booking)
	if err != nil {
		t.Errorf("Booking save failed: %v", err)
	}

	seats, err := db.GetBookedSeats("Dune", "10:00 PM")
	if err != nil {
		t.Errorf("Fetching booked seats failed: %v", err)
	}

	found := false
	for _, s := range seats {
		if s == "A1" || s == "A2" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Saved seats not found in booked list")
	}
}

// ✅ Test JSON marshalling/unmarshalling for ticket
func TestJSONMarshalling(t *testing.T) {
	ticket := models.SavedTicket{
		Filename:  "ticket-Dune-test.json",
		Movie:     "Dune",
		Showtime:  "10:00 PM",
		Seats:     []string{"B4"},
		User:      models.User{Name: "Alice", Email: "a@b.com", Phone: "123"},
		Timestamp: time.Now(),
	}

	data, err := json.Marshal(ticket)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var unmarshalled models.SavedTicket
	err = json.Unmarshal(data, &unmarshalled)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if unmarshalled.Movie != ticket.Movie {
		t.Errorf("Expected movie %s, got %s", ticket.Movie, unmarshalled.Movie)
	}
}

// ✅ Simulate a mock payment processor
func ProcessPayment(user string, amount int) error {
	if user == "fail@example.com" {
		return errors.New("Payment gateway error")
	}
	return nil
}

func TestProcessPayment(t *testing.T) {
	tests := []struct {
		user    string
		amount  int
		wantErr bool
	}{
		{"success@example.com", 500, false},
		{"fail@example.com", 400, true},
	}

	for _, tc := range tests {
		err := ProcessPayment(tc.user, tc.amount)
		if (err != nil) != tc.wantErr {
			t.Errorf("Payment test failed for %s: got err=%v, wantErr=%v", tc.user, err, tc.wantErr)
		}
	}
}

// ✅ Test seat availability logic
func TestRealTimeSeatAvailability(t *testing.T) {
	bookingMu.Lock()
	defer bookingMu.Unlock()

	_ = db.SaveBooking(models.Booking{
		Movie:    "Joker",
		Showtime: "7:00 PM",
		Seats:    []string{"C1", "C2"},
	})

	booked, err := db.GetBookedSeats("Joker", "7:00 PM")
	if err != nil {
		t.Fatal(err)
	}

	m := make(map[string]bool)
	for _, s := range booked {
		m[s] = true
	}

	if !m["C1"] || !m["C2"] {
		t.Error("Real-time seat availability check failed")
	}
}
