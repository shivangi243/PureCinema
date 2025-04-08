package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var showtimes = []string{
	"10:00 AM",
	"1:00 PM",
	"4:00 PM",
	"7:00 PM",
	"10:00 PM",
}

func ShowtimeScreen(w fyne.Window, movie string) fyne.CanvasObject {
	showtimeButtons := make([]fyne.CanvasObject, 0)

	for _, time := range showtimes {
		t := time // capture variable
		btn := widget.NewButton(t, func() {
			// Move to seat selection screen
			w.SetContent(SeatSelectionScreen(w, movie, t))
		})
		showtimeButtons = append(showtimeButtons, btn)
	}

	backBtn := widget.NewButton("Back to Movies", func() {
		w.SetContent(MovieListScreen(w))
	})

	return container.NewVBox(
		widget.NewLabelWithStyle(fmt.Sprintf("🎟 Showtimes for %s", movie), fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewVBox(showtimeButtons...),
		backBtn,
	)
}
