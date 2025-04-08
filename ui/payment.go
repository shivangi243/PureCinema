package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// PaymentScreen shows a dummy payment UI before confirming booking
func PaymentScreen(w fyne.Window, movie, showtime string, seats []string, email string, price int) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("\U0001F4B3 Payment", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	methods := []string{"Card", "UPI", "Wallet"}
	methodDropdown := widget.NewSelect(methods, nil)
	methodDropdown.PlaceHolder = "Select Payment Method"

	amountLabel := widget.NewLabelWithStyle(fmt.Sprintf("\U0001F4B0 Total: ₹%d", price), fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	detailsEntry := widget.NewMultiLineEntry()
	detailsEntry.SetPlaceHolder("Enter simulated payment details...")

	payBtn := widget.NewButton("✔️ Pay Now", func() {
		if methodDropdown.Selected == "" {
			dialog.ShowError(fmt.Errorf("please select a payment method"), w)
			return
		}

		// Show success popup before redirecting to confirmation
		dialog.ShowInformation("Payment Successful", "Your payment has been processed successfully.", w)

		// Delay a bit then show confirmation
		go func() {
			seatStr := ""
			for i, s := range seats {
				if i > 0 {
					seatStr += ", "
				}
				seatStr += s
			}
			w.SetContent(ConfirmationScreen(w, movie, showtime, seatStr, email))
			w.Canvas().Refresh(w.Content())
		}()
	})

	cancelBtn := widget.NewButton("❌ Cancel", func() {
		w.SetContent(SeatSelectionScreen(w, movie, showtime))
	})

	return container.NewVBox(
		layout.NewSpacer(),
		title,
		widget.NewLabelWithStyle("\U0001F4C4 Select Payment Method", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		methodDropdown,
		amountLabel,
		detailsEntry,
		container.NewHBox(layout.NewSpacer(), payBtn, cancelBtn, layout.NewSpacer()),
		layout.NewSpacer(),
	)
}
