package main

import "fmt"

func main() {
	type location struct {
		lat, lng float64
	}

	country1 := location{lat: 90.1, lng: 80.0}
	country2 := country1

	country1.lat += 10.0

	fmt.Printf("%+v\n", country1)
	fmt.Printf("%+v\n", country2)
}
