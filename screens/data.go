package screens

import "fyne.io/fyne/v2"

type Screen struct {
	title string
	View  func(w fyne.Window) fyne.CanvasObject
}

var Screens = map[string]Screen{
	"youtube": {"Youtube", youtubeScreen},
}
