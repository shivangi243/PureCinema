package ui

import (
	"cinema/models"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var movies = []models.Movie{
	{"Inception", "2h 28m", "assets/inception.jpg"},
	{"Interstellar", "2h 49m", "assets/interstellar.jpg"},
	{"Avengers: Endgame", "3h 1m", "assets/endgame.jpg"},
	{"The Batman", "2h 56m", "assets/batman.jpg"},
	{"Dune", "2h 35m", "assets/dune.jpg"},
	{"Oppenheimer", "3h 0m", "assets/oppenheimer.jpg"},
	{"Joker", "2h 2m", "assets/joker.jpg"},
	{"Avatar", "2h 42m", "assets/avatar.jpg"},
	{"The Godfather", "2h 55m", "assets/godfather.jpg"},
	{"Tenet", "2h 30m", "assets/tenet.jpg"},
	{"Parasite", "2h 12m", "assets/parasite.jpg"},
	{"Shawshank Redemption", "2h 22m", "assets/shawshank.jpg"},
}

func MovieListScreen(w fyne.Window) fyne.CanvasObject {
	header := canvas.NewText("🎬 Welcome to MovieMunch", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	header.TextSize = 28
	header.Alignment = fyne.TextAlignCenter
	header.TextStyle = fyne.TextStyle{Bold: true}

	grid := container.NewGridWithColumns(4)
	for _, m := range movies {
		grid.Add(createMovieCard(m, w))
	}

	scrollableGrid := container.NewVScroll(grid)
	scrollableGrid.SetMinSize(fyne.NewSize(800, 600))

	logoutBtn := widget.NewButtonWithIcon("Logout", theme.LogoutIcon(), func() {
		w.SetContent(LoginScreen(w))
	})

	return container.NewBorder(header, logoutBtn, nil, nil, scrollableGrid)
}

func createMovieCard(movie models.Movie, w fyne.Window) fyne.CanvasObject {
	var img *canvas.Image
	if _, err := os.Stat(movie.ImagePath); err == nil {
		img = canvas.NewImageFromFile(movie.ImagePath)
	} else {
		img = canvas.NewImageFromResource(theme.WarningIcon())
	}
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.NewSize(160, 220))

	title := widget.NewLabelWithStyle(movie.Title, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	runtime := widget.NewLabelWithStyle(movie.Runtime, fyne.TextAlignCenter, fyne.TextStyle{})

	bookBtn := widget.NewButton("🎟️ Book Now", func() {
		w.SetContent(ShowtimeScreen(w, movie.Title))
	})

	cardContent := container.NewVBox(
		img,
		title,
		runtime,
		bookBtn,
	)

	card := container.NewVBox(cardContent)
	card.Resize(fyne.NewSize(180, 300))
	return card
}
