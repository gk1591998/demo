package main

import "fmt"

func average(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {

	numbers := []float64{1, 2, 3, 4, 5}
	fmt.Printf("The average is: %.2f\n", average(numbers))
}
