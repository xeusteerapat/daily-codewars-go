/*
Your task is to remove all duplicate words from a string, leaving only single (first) words entries.

Example:

Input:

'alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta'

Output:

'alpha beta gamma delta'
*/

package main

import (
	"fmt"
	"strings"
)

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}

	return false
}

func RemoveDuplicateWords(str string) string {
	strSlice := strings.Split(str, " ")
	uniqueWords := []string{}

	for _, word := range strSlice {
		if contains(uniqueWords, word) {
			continue
		}

		uniqueWords = append(uniqueWords, word)
	}

	return strings.Join(uniqueWords, " ")
}

func main() {
	// ans := RemoveDuplicateWords("alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta")
	ans := RemoveDuplicateWords("Dy ZAZ Awj Cse kuSTUgbI lSQi UOmEEG")

	fmt.Println(ans)
}
