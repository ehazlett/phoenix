package plugins

import (
	"github.com/Sirupsen/logrus"
	"github.com/ehazlett/phoenix"
)

type (
	ExamplePlugin struct {
		name    string
		version string
		author  string
	}
)

func Example() Plugin {
	plugin := ExamplePlugin{
		name:    "example",
		version: "0.1",
		author:  "ehazlett",
	}
	return plugin
}

func (plugin ExamplePlugin) Name() string {
	return plugin.name
}

func (plugin ExamplePlugin) Version() string {
	return plugin.version
}

func (plugin ExamplePlugin) Author() string {
	return plugin.author
}

func (plugin ExamplePlugin) Handle(message *phoenix.Message) (string, error) {
	logger.WithFields(logrus.Fields{
		"token":       message.Token,
		"teamId":      message.TeamId,
		"channelId":   message.ChannelId,
		"channelName": message.ChannelName,
		"timestamp":   message.Timestamp,
		"userId":      message.UserId,
		"username":    message.Username,
		"text":        message.Text,
		"triggerWord": message.TriggerWord,
	}).Info("Example Plugin Handler")
	return "example plugin", nil
}
