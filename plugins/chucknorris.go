package plugins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ehazlett/phoenix"
)

var (
	ICNDB_URL = "http://api.icndb.com/jokes/random"
)

type (
	ChuckNorrisPlugin struct {
		name        string
		version     string
		author      string
		description string
		apiKey      string
	}

	IcndbResponse struct {
		Type  string     `json:"type"`
		Value IcndbValue `json:"value"`
	}

	IcndbValue struct {
		Id   int    `json:"id"`
		Joke string `json:"joke"`
	}
)

func ChuckNorris() Plugin {
	plugin := ChuckNorrisPlugin{
		name:        "chucknorris",
		version:     "0.1",
		author:      "ehazlett",
		description: "random chuck norris awesomeness",
	}
	return plugin
}

func (plugin ChuckNorrisPlugin) Name() string {
	return plugin.name
}

func (plugin ChuckNorrisPlugin) Version() string {
	return plugin.version
}

func (plugin ChuckNorrisPlugin) Author() string {
	return plugin.author
}

func (plugin ChuckNorrisPlugin) Description() string {
	return plugin.description
}

func (plugin ChuckNorrisPlugin) Handle(message *phoenix.Message) (string, error) {
	resp, err := plugin.randomChuck()
	if err != nil {
		return "", err
	}
	return resp.Value.Joke, nil
}

func (plugin ChuckNorrisPlugin) randomChuck() (*IcndbResponse, error) {
	url := fmt.Sprintf("%s", ICNDB_URL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var icndbResponse IcndbResponse
	if err := json.NewDecoder(resp.Body).Decode(&icndbResponse); err != nil {
		return nil, err
	}
	return &icndbResponse, nil
}
