package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ehazlett/phoenix"
)

var (
	LEBOWSKI_URL = "http://lebowski.me/api/quotes/random"
)

type (
	LebowskiPlugin struct {
		name    string
		version string
		author  string
		apiKey  string
	}

	LebowskiResponse struct {
		Quote LebowskiQuote `json:"quote"`
	}

	LebowskiQuote struct {
		Id    int            `json:"id"`
		Lines []LebowskiLine `json:"lines"`
	}

	LebowskiLine struct {
		Id        int               `json:"id"`
		Text      string            `json:"text"`
		Character LebowskiCharacter `json:"character"`
	}

	LebowskiCharacter struct {
		Name string `json:"name"`
	}
)

func Lebowski() Plugin {
	plugin := LebowskiPlugin{
		name:    "lebowski",
		version: "0.1",
		author:  "ehazlett",
	}
	return plugin
}

func (plugin LebowskiPlugin) Name() string {
	return plugin.name
}

func (plugin LebowskiPlugin) Version() string {
	return plugin.version
}

func (plugin LebowskiPlugin) Author() string {
	return plugin.author
}

func (plugin LebowskiPlugin) Handle(message *phoenix.Message) (string, error) {
	resp, err := plugin.random()
	if err != nil {
		return "", err
	}
	data := ""
	for _, line := range resp.Quote.Lines {
		data += fmt.Sprintf("> *%s*: %s\n", line.Character.Name, line.Text)
	}
	return data, nil
}

func (plugin LebowskiPlugin) random() (*LebowskiResponse, error) {
	url := fmt.Sprintf("%s", LEBOWSKI_URL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var lebowskiResponse LebowskiResponse
	if err := json.NewDecoder(resp.Body).Decode(&lebowskiResponse); err != nil {
		return nil, err
	}
	return &lebowskiResponse, nil
}
