package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetGeocode(adress string) (latitude float32, longitude float32) {

	type Coord struct {
		Lat float32 `json:"latitude"`
		Lon float32 `json:"longitude"`
	}

	type Data struct {
		Data []Coord `json:"data"`
	}

	const GEO_API_KEY = "YOUR_API_KEY&query="
	var GEO_URL = "http://api.positionstack.com/v1/forward?access_key="
	response, err := http.Get((GEO_URL + GEO_API_KEY + adress))
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	var results Data
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		panic(err.Error())
	}

	latitude = results.Data[0].Lat
	longitude = results.Data[0].Lon

	fmt.Println(results.Data[0])

	return latitude, longitude
}
