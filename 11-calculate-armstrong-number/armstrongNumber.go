package main

type MyInt int

func (n *MyInt) IsArmstrong() bool {
	digits := n.SplitToDigits()

	var sum int

	for _, digit := range digits {
		sum += digit * digit * digit
	}

	return sum == int(*n)
}

func (n *MyInt) SplitToDigits() []int {
	digits := make([]int, 0)

	number := int(*n)
	for i := 100; i > 0; i /= 10 {
		digits = append(digits, number/i)
		number = number % i
	}

	return digits
}
