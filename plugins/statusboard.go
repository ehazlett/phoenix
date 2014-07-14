package plugins

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/ehazlett/phoenix"
)

type (
	StatusBoardPlugin struct {
		name        string
		version     string
		author      string
		description string
		status      map[string]string
	}
)

func StatusBoard() Plugin {
	plugin := StatusBoardPlugin{
		name:        "statusboard",
		version:     "0.1",
		author:      "ehazlett",
		description: "records and reports user status",
		status:      make(map[string]string),
	}
	return plugin
}

func (plugin StatusBoardPlugin) Name() string {
	return plugin.name
}

func (plugin StatusBoardPlugin) Version() string {
	return plugin.version
}

func (plugin StatusBoardPlugin) Author() string {
	return plugin.author
}

func (plugin StatusBoardPlugin) Description() string {
	return plugin.description
}

func (plugin StatusBoardPlugin) setUserStatus(username, status string, timestamp time.Time) {
	logger.WithFields(logrus.Fields{
		"username": username,
		"status":   status,
	}).Info("Updating Status")
	plugin.status[username] = fmt.Sprintf("[%s] %s", timestamp, status)
}

func (plugin StatusBoardPlugin) getUserStatus(username string) string {
	status := fmt.Sprintf("%s: ", username)
	if val, ok := plugin.status[username]; ok {
		logger.Info(val)
		status += val
	} else {

		status += "nothing reported"
	}
	return status
}

func (plugin StatusBoardPlugin) getAllUserStatuses() string {
	status := ""
	for k, _ := range plugin.status {
		status += fmt.Sprintf("%s\n", plugin.getUserStatus(k))
	}
	return status
}

func (plugin StatusBoardPlugin) Handle(message *phoenix.Message) (string, error) {
	msgParts := strings.Split(message.Text, " ")
	command := msgParts[0]
	text := strings.Join(msgParts[1:], " ")
	username := message.Username
	switch command {
	case "update":
		plugin.setUserStatus(username, text, message.Timestamp)
	case "user":
		user := msgParts[1]
		return plugin.getUserStatus(user), nil
	case "report":
		return plugin.getAllUserStatuses(), nil
	default:
		logger.WithFields(logrus.Fields{
			"plugin":  "statusboard",
			"command": command,
			"user":    username,
		}).Error("unknown command")
		return "", errors.New("unknown command")
	}
	return "", nil
}
