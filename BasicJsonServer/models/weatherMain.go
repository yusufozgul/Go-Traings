package models

type WeatherMain struct {
	Temp     int `json:"temp"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
	TempMin  int `json:"temp_min"`
	TempMax  int `json:"temp_max"`
}
