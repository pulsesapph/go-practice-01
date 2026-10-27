package main

import "fmt"

func main() {
	first := 20
	second := 25
	third := 35

	sum := first + second + third
	average := float64(sum) / 3

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Среднее: %.2f\n", average)
}
