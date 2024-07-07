package main

import (
	"fmt"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func main() {
	fmt.Println("Go Queue Implementation")

	queue := Queue{}

	flights := []Flight{
		Flight{Origin: "A", Destination: "B", Price: 100},
		Flight{Origin: "C", Destination: "D", Price: 200},
	}

	for _, flight := range flights {
		queue.Push(flight)
	}

	for !queue.IsEmpty() {
		flight, err := queue.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(flight)
	}
}
