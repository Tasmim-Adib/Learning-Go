package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type country struct {
		Name       string
		Population int
	}

	type location struct {
		Lat float64 `json:"latitude"` //alias
		Lng float64 `json:"longitude"`
	}

	shaplaChattor := location{Lat: 90.1, Lng: 80.0}
	bangladesh := country{"Bangladesh", 10000}

	bytes, err := json.Marshal(shaplaChattor)
	bytes, err = json.Marshal(bangladesh)

	exitOnErr(err)
	fmt.Println(string(bytes))

}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
