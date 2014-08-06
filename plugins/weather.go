package plugins

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ehazlett/phoenix"
)

var (
	WEATHER_URL = "http://api.openweathermap.org/data/2.5/find?units=imperial&q="
)

type (
	WeatherPlugin struct {
		name        string
		version     string
		author      string
		description string
		apiKey      string
	}

	WeatherResponse struct {
		Message string `json:"message"`
		List    []List `json:"list"`
	}

	List struct {
		Name    string   `json:"name"`
		Main    MainList `json:"main"`
		Wind    WindList `json:"wind"`
		Sys     SysList  `json:"sys"`
		Weather []Wthr   `json:"weather"`
	}

	Wthr struct {
		Main string `json:"main"`
		Desc string `json:"description"`
	}

	MainList struct {
		Temp float32 `json:"temp"`
	}

	WindList struct {
		Speed float32 `json:"speed"`
	}

	SysList struct {
		Country string `json:"country"`
	}
)

func Weather() Plugin {
	plugin := WeatherPlugin{
		name:        "weather",
		version:     "0.1",
		author:      "mbentley",
		description: "weather plugin",
	}
	return plugin
}

func (plugin WeatherPlugin) Name() string {
	return plugin.name
}

func (plugin WeatherPlugin) Version() string {
	return plugin.version
}

func (plugin WeatherPlugin) Author() string {
	return plugin.author
}

func (plugin WeatherPlugin) Description() string {
	return plugin.description
}

func (plugin WeatherPlugin) Handle(message *phoenix.Message) (string, error) {
	searchText := url.QueryEscape(message.Text)
	if searchText == "" {
		return "", errors.New("you must enter a *full* city name")
	}
	resp, err := plugin.Weather(searchText)
	if err != nil {
		return "", err
	}
	data := ""
	if len(resp.List) == 0 {
		return "", err
	}
	if resp.List[0].Name == "" {
		return "", errors.New("city name not recognized; you must enter a *full* city name")
	}
	data += fmt.Sprintf("> *%s*", resp.List[0].Name)
	data += fmt.Sprintf(", *%s*", resp.List[0].Sys.Country)
	a := int(resp.List[0].Main.Temp)
	data += fmt.Sprintf(" - %dF", a)
	b := int(resp.List[0].Wind.Speed)
	data += fmt.Sprintf(" - %s", resp.List[0].Weather[0].Main)
	data += fmt.Sprintf(" (%s)", resp.List[0].Weather[0].Desc)
	data += fmt.Sprintf(" - Wind: %dmph", b)
	data += "\n"
	return data, nil
}

func (plugin WeatherPlugin) Weather(query string) (*WeatherResponse, error) {
	url := fmt.Sprintf("%s%s", WEATHER_URL, query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, err
	}
	return &weatherResponse, nil
}
