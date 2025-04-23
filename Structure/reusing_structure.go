package main

import "fmt"

func main() {
	type location struct {
		lat, lng float64
	}

	var country1 location
	country1.lat, country1.lng = 90.1, 80.0

	var country2 location
	country2.lat, country2.lng = 110.0, 100.0

	fmt.Println(country1, country2)
}
