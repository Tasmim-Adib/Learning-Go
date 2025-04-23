package main

import "fmt"

type temperature struct {
	high, low float64
}

type place struct {
	lat, lng float64
}

type report struct {
	sol         int
	temperature temperature
	place       place
}
type celsius float64

func (t temperature) averageTemperature() celsius {
	return celsius((t.high + t.low) / 2)
}

func (r report) averageTemperature() celsius {
	return r.temperature.averageTemperature()
}
func main() {
	bangladesh := place{12.9, 32.2}
	temp := temperature{45, 12}
	report := report{10, temp, bangladesh}
	fmt.Println(report)
	fmt.Printf("Temperature Average : %.2f\n", report.temperature.averageTemperature())
	fmt.Printf("Temperature Average via report: %.2f\n", report.averageTemperature())
}
