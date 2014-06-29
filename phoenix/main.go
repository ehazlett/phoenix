package main

import (
	"flag"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/ehazlett/phoenix/plugins"
)

var (
	pluginStr  string
	pluginList []string
	listenAddr string
	logger     = logrus.New()
	version    = "0.1"
)

func init() {
	flag.StringVar(&pluginStr, "plugins", "", "list of enabled plugins (comma separated)")
	flag.StringVar(&listenAddr, "listen", ":8080", "listen address")
	flag.Parse()
}

func main() {
	logger.Infof("phoenix v%s", version)
	pluginList = strings.Split(pluginStr, ",")
	logger.WithFields(logrus.Fields{
		"names": pluginList,
	}).Info("enabling plugins")
	manager := plugins.New(pluginList)
	server := NewServer(manager, listenAddr)
	server.Run()
}
