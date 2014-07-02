package plugins

import (
	"encoding/base64"

	"github.com/ehazlett/phoenix"
)

type (
	Base64Plugin struct {
		name        string
		version     string
		author      string
		description string
	}
)

func Base64() Plugin {
	plugin := Base64Plugin{
		name:        "base64",
		version:     "0.1",
		author:      "ehazlett",
		description: "encodes input as base64",
	}
	return plugin
}

func (plugin Base64Plugin) Name() string {
	return plugin.name
}

func (plugin Base64Plugin) Version() string {
	return plugin.version
}

func (plugin Base64Plugin) Author() string {
	return plugin.author
}

func (plugin Base64Plugin) Description() string {
	return plugin.description
}

func (plugin Base64Plugin) Handle(message *phoenix.Message) (string, error) {
	data := []byte(message.Text)
	res := base64.StdEncoding.EncodeToString(data)
	return res, nil
}
