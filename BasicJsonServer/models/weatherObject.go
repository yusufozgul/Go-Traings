package models

type weatherObject struct {
	coord   coordinat
	weather []weather
	main    weatherMain
	id      int
	name    string
}
