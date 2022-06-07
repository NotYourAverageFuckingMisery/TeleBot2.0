package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetWeather(lat float32, lon float32) (WeatherStatus string) {

	type Weather struct {
		Main string `json:"main"`
		Desc string `json:"description"`
	}

	type Temperature struct {
		CurrentTemp float32 `json:"temp"`
	}

	type Wind struct {
		CurrentWind float32 `json:"speed"`
	}

	type Resp struct {
		Weath []Weather   `json:"weather"`
		Temp  Temperature `json:"main"`
		W     Wind        `json:"wind"`
	}

	const WEATHER_API_KEY = "YOUR_API_KEY"
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + fmt.Sprintf("%f", lat) + "&lon=" + fmt.Sprintf("%f", lon) + "&appid=" + WEATHER_API_KEY)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	var results Resp
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(results)

	// they send temp in kelvin
	CelsiusTemp := (results.Temp.CurrentTemp - 273.15)
	WeatherStatus = fmt.Sprintf("Current weather: %s, %s. Wind speed: %fM/S. Temperature: %f°C", results.Weath[0].Main, results.Weath[0].Desc, results.W.CurrentWind, CelsiusTemp)

	return WeatherStatus
}
