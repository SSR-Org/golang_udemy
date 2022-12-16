package main

import "fmt"

func main() {
	colour := map[string]string{
		"blue":  "BLUE",
		"red":   "RED",
		"green": "GREEN",
	}
	fmt.Println(colour)

	var colour1 map[string]string
	fmt.Println(colour1)

	colour2 := make(map[string]string)
	fmt.Println(colour2)

	delete(colour, "red")
	fmt.Println(colour)

	printMap(colour)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("the colour", key, "is", value)
	}
}
