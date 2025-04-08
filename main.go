package main

import (
	"cinema/db"
	"cinema/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	db.Connect() // Connect to MySQL

	a := app.NewWithID("cinema.booking.app")
	a.Settings().SetTheme(theme.DarkTheme()) // optional dark theme

	w := a.NewWindow("Cinema Booking System")
	w.Resize(fyne.NewSize(400, 500))

	w.SetContent(container.NewMax(
		ui.Router(w),
	))

	w.ShowAndRun()
}
