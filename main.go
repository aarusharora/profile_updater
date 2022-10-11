package main

import (
	"fmt"

	"github.com/aarusharora/profile_updater/src/quote"
	"github.com/aarusharora/profile_updater/src/utils"
	"github.com/aarusharora/profile_updater/src/weather"
)

func main() {
	// get weather for today
	weatherData := "- " + weather.GetData() + "\n"
	fmt.Printf("adding %s to profile.md \n", weatherData)
	// get quote of the day
	quoteData := "- " + quote.GetData() + "\n"
	fmt.Printf("adding %s to profile.md \n", quoteData)
	final_string := weatherData + quoteData
	if utils.AppendToMarkdown("input/profile_markdown.md", "output/README.md", final_string) {
		fmt.Println("updated profile.md")
	}
}
