package main

import "fmt"

func Multiple3And5(number int) int {
	sum := 0

	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println(i)
			sum += i
		}
	}

	return sum
}

func main() {
	a := Multiple3And5(10)

	fmt.Println(a)
}
