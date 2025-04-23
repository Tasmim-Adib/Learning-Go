package main

import "fmt"

func main() {
	type location struct {
		lat, lng float64
	}

	// initialize structure using field-value pairs
	country1 := location{lat: 90.1, lng: 80.5}
	fmt.Println(country1)

	country2 := location{lat: 9.01, lng: 1000.2}
	fmt.Println(country2)

	// print field names with values
	fmt.Printf("%+v\n", country1)

	//initialize structure using only value with orders
	fmt.Println("This struct is initialized only values")
	country1 = location{9.01, 12.90}
	fmt.Printf("%+v\n", country1)
}
