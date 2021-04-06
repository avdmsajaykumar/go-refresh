package main

import (
	"fmt"
	"strings"
)

func StringReverse(input string) string{
	runes := []rune(input)
	for i, j :=0, len(runes) - 1 ; i<j; i,j = i + 1, j - 1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func SentenceReverse(strin string) string {
	strins := strings.Split(strin, " ")
	totalLength := len(strins)
	fmt.Println(totalLength)

	maxLength := 0
	for _, value := range strins {
		if len(value) > maxLength {
			maxLength = len(value)
		}
	}

	fmt.Println(maxLength)
	output := make([]string, 0)
	minLength := 0

	for minLength <= maxLength {
		for i := 0; i < totalLength; i++ {
			word := strins[i]
			word = strings.Trim(word, ",")
			if len(word) == minLength {
				if len(output) == 0 {
					output = append(output, strings.Title(strings.ToLower(word)))
				} else {
					output = append(output, strings.ToLower(word))
				}
			}
		}
		minLength++
	}

	
	return strings.Join(output, " ")
}