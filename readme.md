# Phoenix
Pluggable bot for [Slack](http://slack.com)

# Example

![Image](http://i.imgur.com/4EzXslw.png)

# Usage

Show Help

`docker run -P ehazlett/phoenix -h`

Run with Plugins

`docker run -P -e GIPHY_API_KEY=12345 ehazlett/phoenix -plugins example,giphy`

# Setup

* Run an instance of phoenix (see above) at an internet accessible URL
* Create an [Outgoing Webhook](https://my.slack.com/services/new/outgoing-webhook) in Slack
  * Set Channel (I use "Any")
  * Set trigger words (I use "phx,phoenix")
  * Enter the above URL under "URLs"
  * (optional) customize bot look and feel

# Plugins
To use a plugin, specify the plugin name as the "trigger word" in the Slack message.  For example: `phoenix example foo` (assuming you have setup a word match for `phoenix`).

The following plugins are available:

## Example
Simple example plugin for reference

Name: `example`

## Base64
Base64 encoding

Name: `base64`

## Giphy
[GiphyAPI](https://github.com/giphy/GiphyAPI) search.  Returns Gif links.

Name: `giphy`

Environment Variables

* `GIPHY_API_KEY`: Giphy API Key.  See https://github.com/giphy/GiphyAPI for details.

# Developing Plugins
You can create your own plugins.  Check the `plugins/example.go` for reference.
