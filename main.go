package main

import (
	"fmt"

	"github.com/aarusharora/profile_updater/src/utils"
	"github.com/aarusharora/profile_updater/src/weather"
)

func main() {
	// get weather for today
	weatherData := "- " + weather.GetData()
	fmt.Printf("adding %s to profile.md \n", weatherData)
	// get quote of the day
	if utils.AppendToMarkdown("input/profile_markdown.md", "output/README.md", weatherData) {
		fmt.Println("updated profile.md")
	}
}
