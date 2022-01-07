package main

import "fmt"

func WordsToMarks(s string) int {
	// your code here
	var sum int
	for _, v := range s {
		sum += int(v) - 96
	}

	return sum
}

func main() {
	ans := WordsToMarks("attitude")
	fmt.Println(ans)
}

// Other solution

/*

func WordsToMarks(s string) int {
  count := 0
  for _, i := range s {
     count += int(i) - 'a' + 1;
  }
  return count
}

*/
