package plugins

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ehazlett/phoenix"
)

var (
	COMMAND_LINE_FU_URL = "http://www.commandlinefu.com/commands"
)

type (
	CommandLineFuPlugin struct {
		name        string
		version     string
		author      string
		description string
		apiKey      string
	}

	CommandLineFuResponse struct {
		Id      string `json:"id"`
		Command string `json:"command"`
		Summary string `json:"summary"`
		Url     string `json:"url"`
	}
)

func CommandLineFu() Plugin {
	plugin := CommandLineFuPlugin{
		name:        "commandlinefu",
		version:     "0.1",
		author:      "ehazlett",
		description: "command line usage and description from commandlinefu.com",
	}
	return plugin
}

func (plugin CommandLineFuPlugin) Name() string {
	return plugin.name
}

func (plugin CommandLineFuPlugin) Version() string {
	return plugin.version
}

func (plugin CommandLineFuPlugin) Author() string {
	return plugin.author
}

func (plugin CommandLineFuPlugin) Description() string {
	return plugin.description
}

func (plugin CommandLineFuPlugin) Handle(message *phoenix.Message) (string, error) {
	searchText := url.QueryEscape(message.Text)
	resp, err := plugin.search(searchText)
	if err != nil {
		return "", err
	}
	res := resp[0]
	respText := fmt.Sprintf("`%s`: %s %s", res.Command, res.Summary, res.Url)
	return respText, nil
}

func (plugin CommandLineFuPlugin) search(query string) ([]CommandLineFuResponse, error) {
	data := []byte(query)
	searchId := base64.StdEncoding.EncodeToString(data)
	url := fmt.Sprintf("%s/matching/%s/%s/sort-by-votes/json", COMMAND_LINE_FU_URL, query, searchId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var clfuResponses []CommandLineFuResponse
	if err := json.NewDecoder(resp.Body).Decode(&clfuResponses); err != nil {
		return nil, err
	}
	return clfuResponses, nil
}
