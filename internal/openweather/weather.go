package openweather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response представляет данные ответа от OpenWeatherMap
type Response struct {
	City struct {
		Name string `json:"name"`
	} `json:"city"`
	List []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"list"`
}

// Fetch выполняет запрос к API и выводит прогноз погоды
func Fetch(apiKey, city string) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&cnt=4&units=metric&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Сервер вернул ошибку: %s", resp.Status)
	}

	var data Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}

	fmt.Printf("Прогноз для %s:\n", data.City.Name)
	for i, item := range data.List {
		fmt.Printf("%d. Температура: %.1f°C, Описание: %s\n", i+1, item.Main.Temp, item.Weather[0].Description)
	}
}
