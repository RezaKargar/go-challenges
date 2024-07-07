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
	fmt.Println("Go Stack Implementation")

	stack := Stack{}

	flights := []Flight{
		Flight{Origin: "A", Destination: "B", Price: 100},
		Flight{Origin: "C", Destination: "D", Price: 200},
	}

	for _, flight := range flights {
		stack.Push(flight)
	}

	for !stack.IsEmpty() {
		flight, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(flight)
	}
}
