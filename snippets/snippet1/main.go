package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))
	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())

	// Function example (not a method - no receiver)
	fmt.Printf("Function call - Sum: %d\n", CalculateSum(two2))
	fmt.Printf("Function call - Multiply: %d\n", Multiply(two2))
}

func CalculateSum(tn TwoInts) int {
	return tn.a + tn.b
}

func Multiply(tn TwoInts) int {
	return tn.a * tn.b
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
