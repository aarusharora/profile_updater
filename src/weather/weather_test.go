package weather

import (
	"fmt"
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
	fmt.Println(got.Weather[0])
	if err != nil {
		fmt.Println(err)
	}
}

// func TestGetData(t *testing.T) {
// 	fmt.Println(GetData())
// }
