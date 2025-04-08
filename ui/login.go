package ui

import (
	"cinema/db"
	"cinema/models"
	"cinema/utils" // ✅ use common validation function

	"database/sql"
	"fmt"
	"strings"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func LoginScreen(w fyne.Window) fyne.CanvasObject {
	title := canvas.NewText("🎬 CINEMA BOOKING LOGIN", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	title.TextSize = 26
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("📧 Email")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("🔒 Password")

	statusLabel := widget.NewLabel("")

	loginBtn := widget.NewButtonWithIcon("Login", theme.LoginIcon(), func() {
		email := strings.TrimSpace(emailEntry.Text)
		password := passwordEntry.Text

		if email == "" || password == "" {
			dialog.ShowError(fmt.Errorf("Please fill in all fields."), w)
			return
		}

		var id int
		err := db.Conn.QueryRow("SELECT id FROM users WHERE email = ? AND password = ?", email, password).Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				statusLabel.SetText("❌ Invalid credentials.")
			} else {
				statusLabel.SetText("⚠️ Login error.")
			}
			return
		}

		models.LoggedInEmail = email
		statusLabel.SetText("✅ Login successful!")
		w.SetContent(MovieListScreen(w))
	})

	registerBtn := widget.NewButton("Register", func() {
		w.SetContent(RegisterScreen(w))
	})

	forgotBtn := widget.NewButton("Forgot Password?", func() {
		showForgotPasswordWindow(w)
	})

	form := container.NewVBox(
		title,
		emailEntry,
		passwordEntry,
		loginBtn,
		registerBtn,
		forgotBtn,
		statusLabel,
	)

	formContainer := container.NewCenter(container.NewVBox(form))
	return container.NewMax(formContainer)
}

func showForgotPasswordWindow(parent fyne.Window) {
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter your registered email")

	newPasswordEntry := widget.NewPasswordEntry()
	newPasswordEntry.SetPlaceHolder("Enter new password")

	resetBtn := widget.NewButton("Reset Password", func() {
		email := strings.TrimSpace(emailEntry.Text)
		newPassword := newPasswordEntry.Text

		if email == "" || newPassword == "" {
			dialog.ShowError(fmt.Errorf("Please fill in all fields."), parent)
			return
		}

		// ✅ Use shared validator from utils
		if !utils.IsPasswordStrong(newPassword) {
			dialog.ShowError(fmt.Errorf("Password must be 8+ characters, include a number and special character."), parent)
			return
		}

		res, err := db.Conn.Exec("UPDATE users SET password = ? WHERE email = ?", newPassword, email)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Database error: %v", err), parent)
			return
		}

		rowsAffected, _ := res.RowsAffected()
		if rowsAffected == 0 {
			dialog.ShowInformation("Not Found", "No user found with this email.", parent)
			return
		}

		dialog.ShowInformation("Success", "Password updated successfully. Please log in.", parent)
	})

	pwWin := fyne.CurrentApp().NewWindow("🔐 Reset Password")
	pwWin.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("Forgot Password", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		emailEntry,
		newPasswordEntry,
		resetBtn,
	))
	pwWin.Resize(fyne.NewSize(400, 250))
	pwWin.Show()
}
