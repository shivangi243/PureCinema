package ui

import (
	"cinema/db"
	"cinema/utils"
	"fmt"
	"strings"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func RegisterScreen(w fyne.Window) fyne.CanvasObject {
	// 🔴 Header styled
	title := canvas.NewText("📋 Register New Account", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	title.TextSize = 26
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Full Name")

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email")

	phoneEntry := widget.NewEntry()
	phoneEntry.SetPlaceHolder("Phone Number")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	confirmPasswordEntry := widget.NewPasswordEntry()
	confirmPasswordEntry.SetPlaceHolder("Confirm Password")

	passwordHint := canvas.NewText("🔒 Min 8 chars, include number & special character", color.Gray{Y: 200})
	passwordHint.TextSize = 12

	termsCheckbox := widget.NewCheck("I agree to the Terms and Conditions", nil)

	registerBtn := widget.NewButton("Register", func() {
		name := strings.TrimSpace(nameEntry.Text)
		email := strings.TrimSpace(emailEntry.Text)
		phone := strings.TrimSpace(phoneEntry.Text)
		password := passwordEntry.Text
		confirmPassword := confirmPasswordEntry.Text

		if name == "" || email == "" || phone == "" || password == "" || confirmPassword == "" {
			dialog.ShowError(fmt.Errorf("All fields are required."), w)
			return
		}

		if !termsCheckbox.Checked {
			dialog.ShowError(fmt.Errorf("You must accept the Terms and Conditions."), w)
			return
		}

		if password != confirmPassword {
			dialog.ShowError(fmt.Errorf("Passwords do not match."), w)
			return
		}

		if !utils.IsPasswordStrong(password) {
			dialog.ShowError(fmt.Errorf("Password must be at least 8 characters long, include a number and a special character."), w)
			return
		}

		_, err := db.Conn.Exec("INSERT INTO users (name, email, phone, password) VALUES (?, ?, ?, ?)", name, email, phone, password)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Registration failed: %v", err), w)
			return
		}

		dialog.ShowInformation("Success", "Account created successfully!", w)
		w.SetContent(LoginScreen(w))
	})

	backBtn := widget.NewButton("Back to Login", func() {
		w.SetContent(LoginScreen(w))
	})

	form := container.NewVBox(
		title,
		nameEntry,
		emailEntry,
		phoneEntry,
		passwordEntry,
		confirmPasswordEntry,
		passwordHint, // ℹ️ Visual password rules
		termsCheckbox,
		registerBtn,
		backBtn,
	)

	formContainer := container.NewCenter(container.NewVBox(form))
	return container.NewMax(formContainer)
}
