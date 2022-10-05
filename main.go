package main

import "fmt"

func doStuff() string {
	return "Done"
}

func main() {
	fmt.Println("Hello world " + doStuff())
}
