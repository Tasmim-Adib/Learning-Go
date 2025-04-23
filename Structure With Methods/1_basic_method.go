package main

type Coordinate struct {
	d, m, s float64
	h       rune
}

func (c Coordinate) Decimal() float64 {
	sign := 1.0

	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
