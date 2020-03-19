package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	loader "../loader"
)

func RunServer() {
	http.HandleFunc("/", HandleWeather)

	http.ListenAndServe(":8080", nil)
}

func HandleWeather(w http.ResponseWriter, r *http.Request) {

	weather := loader.LoadWeatherDaha()

	data, error := json.Marshal(weather)
	if error != nil {
		fmt.Println(error)
	}
	w.Write(data)
}
