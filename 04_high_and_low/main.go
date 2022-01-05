/*
HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"

Notes
- All numbers are valid Int32, no need to validate them.
- There will always be at least one number in the input string.
- Output string must be two numbers separated by a single space, and highest number is first.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	// Code here or
	strSlice := strings.Split(in, " ")
	numSlice := []int32{}

	for _, v := range strSlice {
		num, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		numSlice = append(numSlice, int32(num))
	}

	var min int32 = numSlice[0]
	var max int32 = numSlice[0]

	for _, v := range numSlice {

		if max < int32(v) {
			max = int32(v)
		}

		if min > int32(v) {
			min = int32(v)
		}
	}

	minStr := strconv.Itoa(int(min))
	maxStr := strconv.Itoa(int(max))

	return maxStr + " " + minStr
}

func main() {
	ans := HighAndLow("1 2 -3 4 5")

	fmt.Println(ans)
}

// Interesting solution
/*
package kata

import (
  "fmt"
  "sort"
  "strconv"
  "strings"
)

func HighAndLow(in string) string {
  numStrings := strings.Fields(in)
  var nums = []int{}

  for _, i := range numStrings {
    j, _ := strconv.Atoi(i)
    nums = append(nums, j)
  }
  sort.Ints(nums)
  return fmt.Sprintf("%d %d", nums[len(nums)-1], nums[0])
}
*/
