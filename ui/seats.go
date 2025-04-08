package ui

import (
	"cinema/db"
	"cinema/models"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var seatRows = []string{"A", "B", "C", "D", "E"}

const colsPerRow = 10

var tierPrices = map[string]int{
	"A": 500, "B": 400, "C": 400, "D": 300, "E": 300,
}

func SeatSelectionScreen(w fyne.Window, movie string, showtime string) fyne.CanvasObject {
	title := canvas.NewText(fmt.Sprintf("🎟️ %s at %s", movie, showtime), color.White)
	title.TextStyle.Bold = true
	title.Alignment = fyne.TextAlignCenter

	selectedSeats := make(map[string]bool)

	bookedSeats, err := db.GetBookedSeats(movie, showtime)
	bookedMap := make(map[string]bool)
	if err == nil {
		for _, s := range bookedSeats {
			bookedMap[s] = true
		}
	}

	seatGrid := container.NewVBox()

	for _, row := range seatRows {
		price := tierPrices[row]
		label := fmt.Sprintf("🎫 Tier %s - ₹%d", row, price)
		seatGrid.Add(widget.NewLabelWithStyle(label, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))

		rowContainer := container.NewGridWithColumns(colsPerRow)
		for i := 1; i <= colsPerRow; i++ {
			seatID := fmt.Sprintf("%s%d", row, i)
			btn := widget.NewButton(seatID, nil)

			if bookedMap[seatID] {
				btn.Disable()
				btn.Importance = widget.DangerImportance
				btn.SetText(fmt.Sprintf("*%s*", btn.Text)) // Add asterisks to simulate italic text
			} else {
				btn.Importance = widget.LowImportance
				btn.OnTapped = func(id string, b *widget.Button) func() {
					return func() {
						if selectedSeats[id] {
							delete(selectedSeats, id)
							b.Importance = widget.LowImportance
						} else {
							selectedSeats[id] = true
							b.Importance = widget.HighImportance
						}
						b.Refresh()
					}
				}(seatID, btn)
			}
			rowContainer.Add(btn)
		}
		seatGrid.Add(rowContainer)
	}

	screenLabel := canvas.NewText("🎬 Screen This Way", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	screenLabel.Alignment = fyne.TextAlignCenter
	screenLabel.TextStyle.Bold = true

	bookBtn := widget.NewButton("✅ Book Now", func() {
		if len(selectedSeats) == 0 {
			dialog.ShowError(fmt.Errorf("⚠️ Please select at least one seat."), w)
			return
		}

		// Convert map to slice
		var seatList []string
		for seat := range selectedSeats {
			seatList = append(seatList, seat)
		}

		// Example static pricing (you can enhance this per tier later)
		price := len(seatList) * 300 // ₹300 per seat (adjust as needed)

		// Redirect to Payment screen
		w.SetContent(PaymentScreen(w, movie, showtime, seatList, models.LoggedInEmail, price))
	})

	backBtn := widget.NewButton("Back to Showtimes", func() {
		w.SetContent(ShowtimeScreen(w, movie))
	})

	return container.NewMax(
		container.NewBorder(title, container.NewVBox(screenLabel, bookBtn, backBtn), nil, nil, container.NewVScroll(seatGrid)),
	)
}
