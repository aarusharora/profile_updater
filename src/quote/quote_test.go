package quote

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetQuoteOfTheDay(t *testing.T) {
	got, err := GetQuoteOfTheDay()
	if err != nil {
		fmt.Println(err)
		t.Errorf("%v", err)
	} else {
		fmt.Println(got)
	}
}

func TestGetData(t *testing.T) {
	got := GetData()
	fmt.Println(got)
	if !strings.Contains(got, "Quote of the day") {
		t.Errorf(`string %s does not have "Quote of the day" `, got)
	}
}
