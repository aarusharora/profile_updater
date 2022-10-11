package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aarusharora/profile_updater/src/utils"
)

var apiKey string
var dataUnvailable string = "- The weather data is currently not available."

type WeatherDescription struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherTemp struct {
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
}

type WeatherData struct {
	Weather []WeatherDescription `json:"weather"`
	Main    WeatherTemp          `json:"main"`
}

type GeolocationData struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// GetLatLong will get the latitute and longitude from the city,state and country
func GetLatLong(city, state, country string) (GeolocationData, error) {
	url := fmt.Sprintf(
		"http://api.openweathermap.org/geo/1.0/direct?q=%s,%s,%s&appid=%s",
		city,
		state,
		country,
		apiKey)
	// fmt.Println(url)
	resp, err := utils.MakeGetRequest(url)
	if err != nil {
		return GeolocationData{}, err
	}
	//Got response
	// fmt.Println(resp)
	var location []GeolocationData
	if err := json.Unmarshal([]byte(resp), &location); err != nil {
		return GeolocationData{}, err
	}
	// fmt.Printf("%f %f", location[0].Latitude, location[0].Longitude)
	return location[0], nil
}

// GetWeatherData will get the weather data from latitute and longitude
func GetWeatherData(location GeolocationData) (WeatherData, error) {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%s&units=metric",
		location.Latitude,
		location.Longitude,
		apiKey)
	resp, err := utils.MakeGetRequest(url)
	if err != nil {
		return WeatherData{}, err
	}
	// fmt.Println(resp)
	var currentWeather WeatherData
	if err := json.Unmarshal([]byte(resp), &currentWeather); err != nil {
		return WeatherData{}, nil
	}
	return currentWeather, nil
}

func GetData() string {
	// get env variables - city, state, country
	city_state_country, ok := os.LookupEnv("CITY_STATE_COUNTRY")
	if !ok {
		fmt.Println("CITY_STATE_COUNTRY not set")
		return dataUnvailable
	}
	var city, state, country string
	utils.UnpackStringSlice(strings.Split(city_state_country, ","), &city, &state, &country)
	// make request
	location, err := GetLatLong(city, state, country)
	if err != nil {
		fmt.Println(err)
		return dataUnvailable
	}
	currentWeather, err := GetWeatherData(location)
	if err != nil {
		fmt.Println(err)
		return dataUnvailable
	}
	// process data
	icon := fmt.Sprintf("![](https://openweathermap.org/img/wn/%s.png)", currentWeather.Weather[0].Icon)
	// fmt.Print(icon)
	final_string := fmt.Sprintf(
		"- The weather for %s is %s %s with temperature from %v℃ - %v℃",
		city, currentWeather.Weather[0].Main, icon, currentWeather.Main.TempMin, currentWeather.Main.TempMax)
	// return result
	return final_string
}

func init() {
	var ok bool
	// fmt.Println(os.Environ())
	if apiKey, ok = os.LookupEnv("API_KEY"); !ok {
		fmt.Println("API_KEY not set")
		apiKey = ""
	}
	// fmt.Printf("got %s api key\n", apiKey)
}
