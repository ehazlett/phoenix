package plugins

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ehazlett/phoenix"
)

var (
	HN_URL = "https://news.ycombinator.com/rss"
)

type (
	HackernewsPlugin struct {
		name    string
		version string
		author  string
		apiKey  string
	}

	HackernewsChannel struct {
		XMLName xml.Name        `xml:"rss"`
		Items   HackernewsItems `xml:"channel"`
	}
	HackernewsItems struct {
		XMLName  xml.Name         `xml:"channel"`
		ItemList []HackernewsItem `xml:"item"`
	}
	HackernewsItem struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
	}
)

func Hackernews() Plugin {
	plugin := HackernewsPlugin{
		name:    "hn",
		version: "0.1",
		author:  "ehazlett",
	}
	return plugin
}

func (plugin HackernewsPlugin) Name() string {
	return plugin.name
}

func (plugin HackernewsPlugin) Version() string {
	return plugin.version
}

func (plugin HackernewsPlugin) Author() string {
	return plugin.author
}

func (plugin HackernewsPlugin) Handle(message *phoenix.Message) (string, error) {
	resp, err := plugin.getLatest()
	if err != nil {
		return "", err
	}
	data := ""
	items := resp.Items.ItemList[:9]
	for _, r := range items {
		data += fmt.Sprintf("%s %s\n", r.Title, r.Link)
	}
	return data, nil
}

func (plugin HackernewsPlugin) getLatest() (*HackernewsChannel, error) {
	url := fmt.Sprintf("%s", HN_URL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var channel HackernewsChannel
	if err := xml.NewDecoder(resp.Body).Decode(&channel); err != nil {
		return nil, err
	}
	return &channel, nil
}
