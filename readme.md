# Phoenix
Pluggable bot for [Slack](http://slack.com)

# Usage

# Plugins
To use a plugin, specify the plugin name as the "trigger word" in the Slack message.  For example: `phoenix example foo` (assuming you have setup a word match for `phoenix`).

The following plugins are available:

* `example`: Simple example plugin for reference
* `giphy`: [GiphyAPI](https://github.com/giphy/GiphyAPI) search (returns gif links)

# Developing Plugins
You can create your own plugins.  Check the `plugins/example.go` for reference.
