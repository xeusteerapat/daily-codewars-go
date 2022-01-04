package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	mynumbers := generateSlice(10)
	fmt.Println("unsorted:", mynumbers)

	ans := SortNumbers(mynumbers)
	fmt.Println("sorted:", ans)
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func SortNumbers(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	left, right := 0, len(numbers)-1

	pivot := rand.Int() % len(numbers)

	numbers[pivot], numbers[right] = numbers[right], numbers[pivot]

	for i := range numbers {
		if numbers[i] < numbers[right] {
			numbers[left], numbers[i] = numbers[i], numbers[left]
			left++
		}
	}

	numbers[left], numbers[right] = numbers[right], numbers[left]

	SortNumbers(numbers[:left])
	SortNumbers(numbers[left+1:])

	return numbers
}
