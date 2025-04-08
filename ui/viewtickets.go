package ui

import (
	"cinema/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ViewTicketsScreen(w fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("📂 Saved Ticket Receipts", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Read current directory for files starting with ticket- and ending with .json
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return widget.NewLabel("Failed to list files.")
	}

	var jsonFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "ticket-") && strings.HasSuffix(file.Name(), ".json") {
			jsonFiles = append(jsonFiles, file.Name())
		}
	}

	if len(jsonFiles) == 0 {
		return container.NewVBox(title, widget.NewLabel("No saved JSON tickets found."))
	}

	list := container.NewVBox()
	for _, file := range jsonFiles {
		f := file // capture loop variable
		btn := widget.NewButton(f, func() {
			data, err := os.ReadFile(f)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to read file: %v", err), w)
				return
			}

			var ticket utils.SavedTicket
			err = json.Unmarshal(data, &ticket)
			if err != nil {
				dialog.ShowError(fmt.Errorf("failed to parse JSON: %v", err), w)
				return
			}

			details := fmt.Sprintf("🎟️ %s\n\nMovie: %s\nShowtime: %s\nSeats: %s\n\n👤 %s\n📧 %s\n📱 %s\n\nDate: %s",
				ticket.Filename, ticket.Movie, ticket.Showtime, strings.Join(ticket.Seats, ", "),
				ticket.User.Name, ticket.User.Email, ticket.User.Phone, ticket.Timestamp.Format("02-Jan-2006 15:04"))

			dialog.ShowInformation("Ticket Info", details, w)
		})
		list.Add(btn)
	}

	back := widget.NewButton("Back to Movies", func() {
		w.SetContent(MovieListScreen(w))
	})

	return container.NewVBox(title, list, back)
}
