package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf754",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
