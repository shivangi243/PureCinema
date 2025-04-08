package db

import (
	"cinema/models"
	"fmt"
)

// Fetch user details by email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	row := Conn.QueryRow("SELECT name, email, phone FROM users WHERE email = ?", email)

	err := row.Scan(&user.Name, &user.Email, &user.Phone)
	if err != nil {
		return user, fmt.Errorf("user fetch failed: %v", err)
	}

	return user, nil
}

// ✅ Final GetBookedSeats function - fetch booked seat IDs as []string
func GetBookedSeats(movie, showtime string) ([]string, error) {
	rows, err := Conn.Query("SELECT seats FROM bookings WHERE movie = ? AND showtime = ?", movie, showtime)
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var seats []string
	for rows.Next() {
		var seat string
		if err := rows.Scan(&seat); err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}
	return seats, nil
}
