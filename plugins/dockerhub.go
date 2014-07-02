package plugins

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ehazlett/phoenix"
)

var (
	HUB_URL = "http://registry.hub.docker.com"
)

type (
	DockerHubPlugin struct {
		name    string
		version string
		author  string
	}

	DockerHubResponse struct {
		Query        string                  `json:"query"`
		NumOfResults int                     `json:"num_results"`
		Results      []DockerHubSearchResult `json:"results"`
	}

	DockerHubSearchResult struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

func DockerHub() Plugin {
	plugin := DockerHubPlugin{
		name:    "dockerhub",
		version: "0.1",
		author:  "ehazlett",
	}
	return plugin
}

func (plugin DockerHubPlugin) Name() string {
	return plugin.name
}

func (plugin DockerHubPlugin) Version() string {
	return plugin.version
}

func (plugin DockerHubPlugin) Author() string {
	return plugin.author
}

func (plugin DockerHubPlugin) Handle(message *phoenix.Message) (string, error) {
	searchText := url.QueryEscape(message.Text)
	if searchText == "" {
		return "", errors.New("you must enter a search term")
	}
	resp, err := plugin.search(searchText)
	if err != nil {
		return "", err
	}
	data := ""
	res := resp.Results
	// limit to 25 if more
	if len(res) > 25 {
		res = res[:24]
	}
	for _, r := range res {
		data += fmt.Sprintf("> %s", r.Name)
		if r.Description != "" {
			data += fmt.Sprintf(": %s", r.Description)
		}
		data += "\n"
	}
	return data, nil
}

func (plugin DockerHubPlugin) search(query string) (*DockerHubResponse, error) {
	url := fmt.Sprintf("%s/v1/search?q=%s", HUB_URL, query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var hubResponse DockerHubResponse
	if err := json.NewDecoder(resp.Body).Decode(&hubResponse); err != nil {
		return nil, err
	}
	return &hubResponse, nil
}
