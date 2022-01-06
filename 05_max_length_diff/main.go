/*
You are given two arrays a1 and a2 of strings.
Each string is composed with letters from a to z.
Let x be any string in the first array and y be any string in the second array.

Find max(abs(length(x) − length(y)))

If a1 and/or a2 are empty return -1 in each language except in Haskell (F#)
where you will return Nothing (None).

Example:
a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
mxdiflg(a1, a2) --> 13
Bash note:
- input : 2 strings with substrings separated by ,
- output: number as a string
*/

package main

import (
	"fmt"
	"math"
)

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	var min1 int = len(a1[0])
	var min2 int = len(a2[0])
	var max1 int
	var max2 int

	for _, v := range a1 {
		if max1 < len(v) {
			max1 = len(v)
		}

		if min1 > len(v) {
			min1 = len(v)
		}
	}

	for _, v := range a2 {
		if max2 < len(v) {
			max2 = len(v)
		}

		if min2 > len(v) {
			min2 = len(v)
		}
	}

	possibleAnswer1 := int(math.Abs(float64(max1) - float64(min2)))
	possibleAnswer2 := int(math.Abs(float64(max2) - float64(min1)))

	if possibleAnswer1 > possibleAnswer2 {
		return possibleAnswer1
	}

	return possibleAnswer2
}

func main() {
	var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	var s2 = []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}

	s1 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	s2 = []string{"bbbaaayddqbbrrrv"}
	ans := MxDifLg(s1, s2)
	fmt.Println(ans)
}
