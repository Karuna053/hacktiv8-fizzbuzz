package main

import (
	"fmt"
)

func main() {
	var countUpTo int
	fmt.Println("How much number are you putting in for this FizzBuzz program?")
	fmt.Scanln(&countUpTo)

	for i := 1; i <= countUpTo; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}

	fmt.Println("Press Enter to exit program!")
	fmt.Scanln() // Wait for input
}
