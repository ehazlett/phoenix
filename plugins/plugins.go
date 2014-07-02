package plugins

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/ehazlett/phoenix"
)

var (
	logger = logrus.New()
)

type (
	Plugin interface {
		Handle(*phoenix.Message) (string, error)
		Name() string
		Version() string
		Author() string
		Description() string
	}

	Manager struct {
		plugins        map[string]Plugin
		enabledPlugins []string
	}
)

func New(enabledPlugins []string) *Manager {
	// initialize plugins
	plugins := make(map[string]Plugin)
	// plugins
	plugins["base64"] = Base64()
	plugins["chucknorris"] = ChuckNorris()
	plugins["commandlinefu"] = CommandLineFu()
	plugins["dockerhub"] = DockerHub()
	plugins["example"] = Example()
	plugins["giphy"] = Giphy()
	plugins["hn"] = Hackernews()
	plugins["lebowski"] = Lebowski()
	// manager
	manager := &Manager{
		plugins:        plugins,
		enabledPlugins: enabledPlugins,
	}
	return manager
}

func (manager *Manager) ShowPluginList() string {
	data := "enabled plugins: \n"
	for _, plugin := range manager.Plugins() {
		data += fmt.Sprintf("    %s %s (%s): %s\n", plugin.Name(), plugin.Version(), plugin.Author(), plugin.Description())
	}
	return data
}

func (manager *Manager) Plugins() map[string]Plugin {
	return manager.plugins
}

func (manager *Manager) EnabledPlugins() []string {
	return manager.enabledPlugins
}

func (manager *Manager) Handle(msg *phoenix.Message) string {
	resp := "unknown plugin"
	if msg.PluginName != "" {
		// handle "special" plugins
		switch msg.PluginName {
		case "info":
			return manager.ShowPluginList()
		default:
			// check for enabled plugin
			for _, p := range manager.enabledPlugins {
				// check if plugin matches trigger
				if msg.PluginName == p {
					resp = manager.runPlugin(p, msg)
				}
			}
		}
	}
	return resp
}

func (manager *Manager) runPlugin(pluginName string, message *phoenix.Message) string {
	resp := ""
	for _, plugin := range manager.plugins {
		// if enabled plugin found, execute
		if plugin.Name() == pluginName {
			logger.WithFields(logrus.Fields{
				"name":    plugin.Name(),
				"version": plugin.Version(),
				"author":  plugin.Author(),
				"text":    message.Text,
			}).Infof("running plugin")
			r, err := plugin.Handle(message)
			if err != nil {
				logMsg := fmt.Sprintf("error handling message: %s", err)
				logger.Errorf(logMsg)
				r = logMsg
			}
			resp = r
			break
		}
	}
	return resp
}
