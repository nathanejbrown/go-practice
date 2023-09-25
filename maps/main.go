package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red": "#F00",
		"green": "#4b7f45",
		"white": "FFF",
	}

	// colors["white"] = "#FFF"

	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("color: ", color, "hex: ", hex)
	}
}