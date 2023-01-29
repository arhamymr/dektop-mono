package apis

import (
	"desktop-mono/configs"
	debugs "desktop-mono/debug-tools"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var maxResults = flag.Int64("max-results", 25, "Max YouTube results")

type SearchResult struct {
	Title       string `json:"title"`
	ChannelId   string `json:"ChannelId"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

func SearchVideos(keyword string) []SearchResult {
	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: configs.GET("GCP_API_KEY")},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	snippet := []string{"id", "snippet"}

	call := service.Search.List(snippet).
		Q(keyword).
		MaxResults(*maxResults)
	response, err := call.Do()

	if err != nil {
		panic("error" + err.Error())
	}

	var searchResult []SearchResult

	bytearray, err := json.Marshal(response.Items)

	if err != nil {
		panic("failed marshal" + err.Error())
	}

	debugs.PrintPrettyJSON(string(bytearray))

	for _, value := range response.Items {
		loopdata := SearchResult{
			value.Snippet.Title,
			value.Snippet.ChannelId,
			value.Snippet.Description,
			value.Snippet.Thumbnails.Default.Url,
		}
		searchResult = append(searchResult, loopdata)
	}

	return searchResult
}
