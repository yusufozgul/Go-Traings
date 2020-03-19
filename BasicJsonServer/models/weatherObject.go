package models

type WeatherObject struct {
	Coord   Coordinat   `json:"coord"`
	Weather Weather     `json:"weather"`
	Main    WeatherMain `json:"main"`
	ID      int         `json:"id"`
	Name    string      `json:"name"`
}
