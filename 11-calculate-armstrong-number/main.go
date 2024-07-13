package main

import (
	"fmt"
)

func main() {
	fmt.Println("Armstrong Numbers")

	var num1 MyInt = 371
	fmt.Println(num1.IsArmstrong())
}
