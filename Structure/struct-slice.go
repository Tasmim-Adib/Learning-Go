package main

import "fmt"

func main() {
	type location struct {
		name     string
		lat, lng float64
	}

	locations := []location{
		location{"Shahid Minar", 91.1, 80.0},
		location{"Shapla Chattor", 89.1, 100.2},
		location{"Farmgate", 12.3, 90.9},
	}

	for x, location := range locations {
		fmt.Printf("%v: %+v\n", x+1, location)
	}
}
