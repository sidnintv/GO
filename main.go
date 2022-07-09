package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi! I am telegrame bot Trading")
	array := []int{-10, 20, 11, 542, -36, 65, 10, 0}
	minValue := array[0]
	for _, value := range array {
		if minValue < 0 {
			minValue = value
		} else if value < minValue && value > 0 && value%2 == 0 {
			minValue = value
		}
	}
	fmt.Println(minValue)
}

/*
func main() {
for i := 1; i <= 20; i++ {
	if i%3 == 0 && i%5 == 0 {
		fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}*/
