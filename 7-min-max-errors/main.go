package main

import (
	"errors"
	"fmt"
	"math"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func GetMinMax(flights []Flight) (int, int, error) {
	if len(flights) == 0 {
		return 0, 0, errors.New("no flights found")
	}

	minItem, maxItem := math.MaxInt64, math.MinInt64
	for _, flight := range flights {
		if flight.Price > maxItem {
			maxItem = flight.Price
		}
		if flight.Price < minItem {
			minItem = flight.Price
		}
	}
	return minItem, maxItem, nil
}

func main() {
	fmt.Println("Getting the Minimum and Maximum Flight Prices")

	flights := []Flight{
		Flight{Origin: "A", Destination: "B", Price: 100},
		Flight{Origin: "C", Destination: "D", Price: 200},
	}
	minItem, maxItem, err := GetMinMax(flights)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The Minimum and Maximum Flight Prices are:", minItem, maxItem)
}
