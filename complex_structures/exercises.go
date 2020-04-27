package main

import (
	"fmt"
)

func printAge(ages ...int) {
	for _, val := range ages {
		fmt.Println(val)
	}
}

// func getAverage(num1, num2, num3 float32) (average float32) {
// 	average = (num1 + num2 + num3) / 3
// 	return
// }

// Variadic Function
func getAverageOptionOne(numbers ...float32) (average float32) {
	for _, val := range numbers {
		average += val
	}
	return average / float32(len(numbers))
}

// Variadic Function
func getAverageOptionTwo(numbers ...float32) (average float32) {
	for i := 0; i < len(numbers); i++ {
		average += numbers[i]
	}
	return average / float32(len(numbers))
}

func main() {
	printAge(15, 28, 29, 20)
	average := getAverageOptionTwo(11, 2, 2)
	fmt.Println(average)
}
