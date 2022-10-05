package main

import "fmt"

func doStuff() string {
	return "Done"
}

func main() {
	fmt.Println("TEST Hello world " + doStuff())
}
