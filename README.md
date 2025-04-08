# PureCinema

# 🎬 Cinema Ticket Booking System (Go + Fyne + MySQL)

A feature-rich desktop GUI application built in Go using the **Fyne** framework for booking cinema tickets. It ensures real-time seat availability, secure bookings, concurrency-safe operations, and aesthetic UI for a complete user experience.

---

## 📁 Project Structure
```
cinema/
├── main.go
├── db/
│   ├── mysql.go          # MySQL connection
│   ├── booking.go        # Booking logic with mutex concurrency
│   ├── user.go           # User-related DB functions
├── models/
│   ├── booking.go        # Booking model
│   ├── user.go           # User model
│   ├── movie.go          # Movie model
│   ├── ticket.go         # Ticket struct for marshalling
├── ui/
│   ├── login.go
│   ├── register.go
│   ├── movies.go
│   ├── showtime.go
│   ├── seats.go
│   ├── confirmation.go
│   ├── viewtickets.go
│   ├── payment.go
├── utils/
│   ├── email.go          # Send ticket via email
│   ├── ticket_json.go    # Save/load ticket to JSON
│   ├── validation.go     # Password validation
├── tests/
│   ├── booking_test.go
│   ├── payment_test.go
│   ├── availability_test.go
└── assets/               # Movie images
```

---

## ✅ Features
- 🔐 Login / Register with password validation
- 📧 Forgot password + update in DB
- 🎬 Movie grid (12+ films with posters)
- 🪑 Seat selection with real-time availability
- 🔒 Concurrency-safe seat booking using `sync.Mutex`
- 💰 Payment simulation (Card/UPI/Wallet)
- 🧾 JSON receipt + Email confirmation
- 📂 View previous tickets
- 🌗 Red-black themed UI with enhanced spacing and responsiveness

---

## 🧪 Unit Tests
Located in the `tests/` folder:
- `booking_test.go`: Test for successful bookings
- `payment_test.go`: Validate payment logic
- `availability_test.go`: Simulate real-time concurrency seat checks

Run tests with:
```bash
go test ./tests/...
```

---

## 💾 Database Setup
```sql
CREATE DATABASE cinema;
USE cinema;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(15),
    password VARCHAR(100)
);

CREATE TABLE bookings (
    id INT AUTO_INCREMENT PRIMARY KEY,
    movie VARCHAR(100),
    showtime VARCHAR(100),
    seats JSON
);
```

---

## 🛠 How to Run
```bash
cd cinema
go mod tidy
go run main.go
```

Ensure:
- MySQL is running.
- You have configured Gmail App password for sending emails.

---

## 👩‍💻 Technologies Used
- **Go** (Golang)
- **Fyne** (GUI framework)
- **MySQL** (Database)
- **Gmail SMTP** (Email receipts)
- **sync.Mutex** for concurrency
- **JSON marshalling** for ticket storage

---

## 📌 Authors & Credits
- Developed by: Shivangi Agarwal
- Inspired by modern ticket booking systems like BookMyShow.

---

## 📦 License
This project is for academic/educational use only. Modify and extend freely.

---

Enjoy booking with **MovieMunch** 🍿

