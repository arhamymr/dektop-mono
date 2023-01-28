package screens

import (
	"fmt"
	"image/color"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kkdai/youtube/v2"
)

func ExamplePlaylist() {
	playlistID := "PLQZgI7en5XEgM0L1_ZcKmEzxW1sCOVZwP"
	client := youtube.Client{}

	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		panic(err)
	}

	/* ----- Enumerating playlist videos ----- */
	header := fmt.Sprintf("Playlist %s by %s", playlist.Title, playlist.Author)
	println(header)
	println(strings.Repeat("=", len(header)) + "\n")

	for k, v := range playlist.Videos {
		fmt.Printf("(%d) %s - '%s'\n", k+1, v.Author, v.Title)
	}

	/* ----- Downloading the 1st video ----- */
	entry := playlist.Videos[0]
	video, err := client.VideoFromPlaylistEntry(entry)
	if err != nil {
		panic(err)
	}
	// Now it's fully loaded.

	fmt.Printf("Downloading %s by '%s'!\n", video.Title, video.Author)

	stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")

	if err != nil {
		panic(err)
	}

	defer file.Close()
	_, err = io.Copy(file, stream)

	if err != nil {
		panic(err)
	}

	println("Downloaded /video.mp4")
}

func ExampleClient() {
	videoID := "BaW_jenozKc"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}

func getData() string {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return sb
}

type listData struct {
	userId int
	id     int
	title  string
	body   string
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func youtubeScreen(_ fyne.Window) fyne.CanvasObject {
	a := app.New()
	w := a.NewWindow("Clock")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}),
			fyne.NewMenuItem("getData", func() {
				fmt.Println(getData())
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	clock := widget.NewLabel("")
	updateTime(clock)

	newWindows := widget.NewButton("Download Video", func() {
		ExampleClient()
	})

	label1 := widget.NewLabel("Label 1")
	value1 := widget.NewLabel("Value")
	label2 := widget.NewLabel("Label 2")
	value2 := widget.NewLabel("Something")
	grid := container.New(layout.NewFormLayout(), label1, value1, label2, value2)

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)

	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3, newWindows, clock)

	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), grid, content, layout.NewSpacer(), text4, layout.NewSpacer())

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	return centered
}
