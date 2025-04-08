package ui

import (
	"cinema/db"
	"cinema/models"
	"cinema/utils"
	"fmt"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ConfirmationScreen(w fyne.Window, movie, showtime, seats, email string) fyne.CanvasObject {
	msg := fmt.Sprintf(
		"🎉 Booking Confirmed!\n\nMovie: %s\nShowtime: %s\nSeats: %s",
		movie, showtime, seats,
	)

	label := widget.NewLabel(msg)

	saveBtn := widget.NewButton("🧾 Download + Email Receipt", func() {
		// Fetch user details from DB
		user, err := db.GetUserByEmail(email)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Could not fetch user details: %v", err), w)
			return
		}

		timestamp := time.Now().Format("2006-01-02-15-04-05")
		filename := fmt.Sprintf("ticket-%s-%s.txt", movie, timestamp)

		content := fmt.Sprintf(`🎟️ CINEMA TICKET

Movie: %s
Showtime: %s
Seats: %s

👤 Name: %s
📧 Email: %s
📱 Phone: %s

Date: %s
`,
			movie, showtime, seats,
			user.Name, user.Email, user.Phone,
			time.Now().Format("02-Jan-2006 15:04"),
		)

		// Save plain text receipt
		err = os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Failed to save receipt: %v", err), w)
			return
		}

		// Save JSON ticket
		ticket := models.SavedTicket{
			Filename:  filename,
			Movie:     movie,
			Showtime:  showtime,
			Seats:     strings.Split(seats, ", "),
			User:      user,
			Timestamp: time.Now(),
		}

		err = utils.SaveTicketAsJSON(ticket)
		if err != nil {
			fmt.Println("JSON save warning:", err)
		}

		// Send email
		err = utils.SendTicket(user.Email, filename)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Email failed: %v", err), w)
			return
		}

		dialog.ShowInformation("Success", "Receipt saved and emailed to:\n"+user.Email, w)
	})

	backBtn := widget.NewButton("Back to Movie List", func() {
		w.SetContent(MovieListScreen(w))
	})

	return container.NewVBox(
		widget.NewLabelWithStyle("✅ Confirmation", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		label,
		saveBtn,
		backBtn,
	)
}
