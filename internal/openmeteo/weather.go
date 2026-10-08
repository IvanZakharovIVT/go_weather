package openmeteo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/innotechdevops/openmeteo"
)

// GeoResult представляет результат геокодирования
type GeoResult struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GeoResponse представляет ответ геокодинга
type GeoResponse struct {
	Results []GeoResult `json:"results"`
}

// Fetch выполняет получение погоды через Open-Meteo API
func Fetch(city string) {
	lat, lon, err := getCoordinates(city)
	if err != nil {
		log.Fatalf("Ошибка получения координат: %v", err)
	}

	fetchWeather(lat, lon, city)
}

// getCoordinates получает координаты города через геокодинг API
func getCoordinates(city string) (float64, float64, error) {
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json", city)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка запроса геокодирования: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("геокодинг вернул ошибку: %s", resp.Status)
	}

	var data GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, 0, fmt.Errorf("ошибка парсинга геокодирования: %w", err)
	}

	if len(data.Results) == 0 {
		return 0, 0, fmt.Errorf("город не найден: %s", city)
	}

	return data.Results[0].Latitude, data.Results[0].Longitude, nil
}

// fetchWeather получает данные о погоде по координатам через библиотеку
func fetchWeather(lat, lon float64, city string) {
	param := openmeteo.Parameter{
		Latitude:       openmeteo.Float32(float32(lat)),
		Longitude:      openmeteo.Float32(float32(lon)),
		CurrentWeather: openmeteo.Bool(true),
	}

	m := openmeteo.New()
	resp, err := m.Execute(param)
	if err != nil {
		log.Fatalf("Ошибка запроса погоды: %v", err)
	}

	fmt.Printf("Погода в %s:\n", city)
	fmt.Printf(resp)
	// fmt.Printf("Температура: %.1f°C\n", resp.CurrentWeather.Temperature)
	// fmt.Printf("Код погоды: %d\n", resp.CurrentWeather.WeatherCode)
	// fmt.Printf("Скорость ветра: %.1f км/ч\n", resp.CurrentWeather.Windspeed)
	// fmt.Printf("Направление ветра: %.1f°\n", resp.CurrentWeather.Winddirection)
}
