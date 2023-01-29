// Package main provides various examples of Fyne API capabilities.
package main

import (
	"desktop-mono/apis"
	"desktop-mono/configs"
	"desktop-mono/layouts"
	"desktop-mono/screens"
	"desktop-mono/utils"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var topWindow fyne.Window

func openTutorial(a fyne.App) {
	wTutor := a.NewWindow("Tutorial")
	wTutor.SetContent(screens.TutorialScreen())
	wTutor.Resize(fyne.NewSize(640, 460))
	wTutor.Show()
}

func main() {
	configs.LoadEnv()
	a := app.NewWithID("Desktop Mono")
	a.SetIcon(theme.FyneLogo())

	// make tray
	utils.MakeTray(a)

	// log lifecycle
	utils.LogLifecycle(a)
	w := a.NewWindow("Desktop Mono")
	topWindow = w
	w.SetMaster()

	// navigation menu

	entry := widget.NewEntry()

	var result []apis.SearchResult

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Search Video", Widget: entry}},
		OnSubmit: func() {
			apis.SearchVideos(entry.Text)
		},
	}

	fmt.Println(result, "result inside")
	// tutorialButton := widget.NewButton("Open new", func() {
	// 	openTutorial(a)
	// })

	w.SetMainMenu(layouts.MakeMenu(a, w))

	// formlayout := container.New(layout.NewFormLayout(), form)
	content := container.New(layout.NewMaxLayout(), form)
	var data = []string{"a", "string", "list"}

	list := container.New(layout.NewMaxLayout(), widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
			o.(*widget.Label).SetText("test")
		}))

	w.SetContent(container.New(layout.NewVBoxLayout(), list, content))
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
