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

## Chuck Norris
Chuck Norris awesomeness

Name: `chucknorris`

## CommandLineFu
Returns commands from commandlinefu.com

Name: `commandlinefu`

## DockerHub
Searches DockerHub (shows top 15 results)

Name: `dockerhub`

## Giphy
[GiphyAPI](https://github.com/giphy/GiphyAPI) search.  Returns Gif links.

Name: `giphy`

Environment Variables

* `GIPHY_API_KEY`: Giphy API Key.  See https://github.com/giphy/GiphyAPI for details.

## Hackernews
Returns top 10 posts from hackernews

Name: `hn`

## Lebowski
Random quotes from The Big Lebowski

Name: `lebowski`

## LMGTFY
Let me Google that for you

Name: `lmgtfy`

## StatusBoard
Record & report user status

Name: `statusboard`

Examples:

* `statusboard update working on phoenix plugins` -- sets status
* `statusboard user ehazlett` -- reports current status for username `ehazlett`
* `statusboard report` -- reports current status for all users

# Developing Plugins
You can create your own plugins.  Check the `plugins/example.go` for reference.

## Testing
To send a test payload to a local dev Phoenix, use the following:

For example, to test the `chucknorris` plugin:

Run phoenix with `./phoenix -plugins chucknorris`

`curl -XPOST 'http://localhost:8080/?token=abcdefg&team_id=1001&channel_id=C12345&channel_name=foo_channel&timestamp=1355517523.000005&user_id=1234&user_name=test&text=phoenix%20chucknorris&trigger_word=phoenix'`
