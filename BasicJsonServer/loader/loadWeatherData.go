package loader

import (
	"encoding/json"
	"fmt"

	models "../models"
	util "../utils"
)

func LoadWeatherDaha() models.WeatherObject {

	bytes, error := util.ReadFile("../json/weather.json")
	if error != nil {
		fmt.Println(error)
	}
	var weather models.WeatherObject

	error = json.Unmarshal(bytes, &weather)
	if error != nil {
		fmt.Println(error)
	}
	return weather
}
