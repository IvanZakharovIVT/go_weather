package main

import (
	"go-weather/internal/openmeteo"
	"go-weather/internal/openweather"
)

func main() {
	const (
		apiKey = "59b4d26d5af213a9a43aeb5b3d82bb74"
		city   = "Voronezh"
	)

	openweather.Fetch(apiKey, city)
	openmeteo.Fetch(city)

}
