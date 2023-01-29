package screens

import (
	"desktop-mono/layouts"
	"desktop-mono/tutorials"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var topWindow fyne.Window

func TutorialScreen() *container.Split {
	a := app.NewWithID("Desktop Mono")
	a.SetIcon(theme.FyneLogo())

	w := a.NewWindow("Desktop Mono")
	topWindow = w

	// navigation menu
	w.SetMainMenu(layouts.MakeMenu(a, w))
	w.SetMaster()

	content := container.NewMax()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")

	setTutorial := func(t tutorials.Tutorial) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(t.Title)
			topWindow = child
			child.SetContent(t.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}

		title.SetText(t.Title)
		intro.SetText(t.Intro)

		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro),
		nil, nil, nil, content,
	)

	split := container.NewHSplit(layouts.MakeNav(setTutorial, true), tutorial)
	split.Offset = 0.2
	w.SetContent(split)
	return split
}
