package main

import "fmt"

func main() {
	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = 4.56
	curiosity.long = -9.89

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}
