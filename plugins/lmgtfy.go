package plugins

import (
	"fmt"
	"net/url"

	"github.com/ehazlett/phoenix"
)

var (
	LMGTFY_URL = "http://lmgtfy.com/?q="
)

type (
	LmgtfyPlugin struct {
		name        string
		version     string
		author      string
		description string
	}
)

func Lmgtfy() Plugin {
	plugin := LmgtfyPlugin{
		name:        "lmgtfy",
		version:     "0.1",
		author:      "ehazlett",
		description: "let me google that for you",
	}
	return plugin
}

func (plugin LmgtfyPlugin) Name() string {
	return plugin.name
}

func (plugin LmgtfyPlugin) Version() string {
	return plugin.version
}

func (plugin LmgtfyPlugin) Author() string {
	return plugin.author
}

func (plugin LmgtfyPlugin) Description() string {
	return plugin.description
}

func (plugin LmgtfyPlugin) Handle(message *phoenix.Message) (string, error) {
	query := url.QueryEscape(message.Text)
	return fmt.Sprintf("%s%s", LMGTFY_URL, query), nil
}
