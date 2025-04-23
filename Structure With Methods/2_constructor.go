package main

import "fmt"

type location struct {
	lat, lng float64
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0

	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, lng coordinate) location {
	return location{lat.decimal(), lng.decimal()}
}

func main() {
	earth := newLocation(coordinate{2, 43, 12.8, 'S'}, coordinate{3, 23, 9.2, 'N'})
	fmt.Println(earth)
}
