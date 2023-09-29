package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	initEnv()

	key := os.Getenv(ENV_KEY)
	api := "https://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=denizli&aqi=no"

	res, err := http.Get(api)

	if err != nil {
		panic(ERROR_1)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(ERROR_1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(ERROR_1)
	}
	var weather Forecast
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(ERROR_1)
	}
	fmt.Println(weather.Location.Country)
}
