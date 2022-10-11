package weather

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetLatLong(t *testing.T) {
	got, err := GetLatLong("Mumbai", "Maharashtra", "India")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("%v %v", got.Latitude, got.Longitude)
	want := GeolocationData{
		Latitude:  19.0785451,
		Longitude: 72.878176,
	}
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestGetWeatherData(t *testing.T) {
	mumbai := GeolocationData{
		Latitude:  19.0785451,
		Longitude: 72.878176,
	}
	got, err := GetWeatherData(mumbai)
	if err != nil {
		fmt.Println(err)
		t.Errorf("got %v", got)
	} else {
		fmt.Println(got.Weather[0])
	}
}

func TestGetData(t *testing.T) {
	got := GetData()
	fmt.Println(got)
	if !strings.Contains(got, "weather") {
		t.Error("string does not contain weather")
	}
}
