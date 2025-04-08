package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Router(w fyne.Window) fyne.CanvasObject {
	// Start with Login screen
	return container.NewMax(
		LoginScreen(w),
	)
}
