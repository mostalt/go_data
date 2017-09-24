package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"black":  "#000",
		"white":  "#fff",
		"orange": "#f40",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex color for", key, "is", value)
	}
}
