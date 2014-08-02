package plugins

import (
	"fmt"
	"net/url"

	"github.com/ehazlett/phoenix"
)

var (
	MANPAGE_URL = "http://manpages.debian.org/cgi-bin/man.cgi?query="
)

type (
	ManpagePlugin struct {
		name        string
		version     string
		author      string
		description string
	}
)

func Manpage() Plugin {
	plugin := ManpagePlugin{
		name:        "manpage",
		version:     "0.1",
		author:      "mbentley",
		description: "debian man page search",
	}
	return plugin
}

func (plugin ManpagePlugin) Name() string {
	return plugin.name
}

func (plugin ManpagePlugin) Version() string {
	return plugin.version
}

func (plugin ManpagePlugin) Author() string {
	return plugin.author
}

func (plugin ManpagePlugin) Description() string {
	return plugin.description
}

func (plugin ManpagePlugin) Handle(message *phoenix.Message) (string, error) {
	query := url.QueryEscape(message.Text)
	return fmt.Sprintf("%s%s", MANPAGE_URL, query), nil
}
