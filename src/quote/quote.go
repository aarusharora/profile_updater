package quote

import (
	"encoding/json"
	"fmt"

	"github.com/aarusharora/profile_updater/src/utils"
)

var dataUnavailable string = "Quote of the day is currently not available."

type QuoteObj struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

// This function will get quote of the day from https://zenquotes.io/api/today API
func GetQuoteOfTheDay() (QuoteObj, error) {
	resp, err := utils.MakeGetRequest("https://zenquotes.io/api/today")
	if err != nil {
		return QuoteObj{}, err
	}
	var quote_of_the_day []QuoteObj
	err = json.Unmarshal([]byte(resp), &quote_of_the_day)
	if err != nil {
		return QuoteObj{}, err
	}
	return quote_of_the_day[0], nil
}

func GetData() string {
	quote_of_the_day, err := GetQuoteOfTheDay()
	if err != nil {
		fmt.Println(err)
		return dataUnavailable
	}
	final_string := fmt.Sprintf(`> Quote of the day:  
	&emsp;"%s"  
	&emsp;&emsp;- %s`, quote_of_the_day.Quote, quote_of_the_day.Author)
	return final_string
}
