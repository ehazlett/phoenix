package plugins

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"

	"github.com/ehazlett/phoenix"
)

var (
	GIPHY_SEARCH_URL = "http://api.giphy.com/v1/gifs/search?q="
)

type (
	GiphyPlugin struct {
		name        string
		version     string
		author      string
		description string
		apiKey      string
	}

	GiphyResponse struct {
		Data []GiphyResult
		Meta GiphyMeta
	}

	GiphyMeta struct {
		Msg    string `json:"msg"`
		Status int    `json:"status"`
	}

	GiphyResult struct {
		Type        string     `json:"type"`
		Id          string     `json:"id"`
		Url         string     `json:"url"`
		BitlyGifUrl string     `json:"bitly_gif_url"`
		BitlyUrl    string     `json:"bitly_url"`
		EmbedUrl    string     `json:"embed_url"`
		Username    string     `json:"username"`
		Source      string     `json:"source"`
		Rating      string     `json:"rating"`
		Images      GiphyImage `json:"images"`
	}

	GiphyImage struct {
		FixedWidth GiphyImageDetails `json:"fixed_width"`
	}

	GiphyImageDetails struct {
		Url    string `json:"url"`
		Width  string `json:"width"`
		Height string `json:"height"`
	}
)

func Giphy() Plugin {
	apiKey := os.Getenv("GIPHY_API_KEY")
	plugin := GiphyPlugin{
		name:        "giphy",
		version:     "0.1",
		author:      "ehazlett",
		description: "searchs giphy for specified gif",
		apiKey:      apiKey,
	}
	return plugin
}

func (plugin GiphyPlugin) Name() string {
	return plugin.name
}

func (plugin GiphyPlugin) Version() string {
	return plugin.version
}

func (plugin GiphyPlugin) Author() string {
	return plugin.author
}

func (plugin GiphyPlugin) Description() string {
	return plugin.description
}

func (plugin GiphyPlugin) Handle(message *phoenix.Message) (string, error) {
	searchText := url.QueryEscape(message.Text)
	resp, err := plugin.searchGiphy(searchText)
	if err != nil {
		return "", err
	}
	img := resp.Data[rand.Intn(len(resp.Data))]
	imgLink := img.Images.FixedWidth.Url
	return imgLink, nil
}

func (plugin GiphyPlugin) searchGiphy(query string) (*GiphyResponse, error) {
	url := fmt.Sprintf("%s%s&api_key=%s", GIPHY_SEARCH_URL, query, plugin.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var giphyResponse GiphyResponse
	if err := json.NewDecoder(resp.Body).Decode(&giphyResponse); err != nil {
		return nil, err
	}
	return &giphyResponse, nil
}
