package main

import (
	"fmt"
)

func main() {
	//create empty map
	//var colors map[string]string

	colors := make(map[int]string)

	// colors := map[string]string{
	// 	"black": "#000",
	// 	"white": "#fff",
	// }

	colors[10] = "#fff"
	delete(colors, 10)

	fmt.Println(colors)
}
