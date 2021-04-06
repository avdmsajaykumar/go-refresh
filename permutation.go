package main

import "fmt"

func Permutation(str string) {
	runes := []rune(str)
	perm(runes, "")
}

func perm(runes []rune, out string) {
	if len(runes) == 0 {
		fmt.Printf("%s ",string(out))
	}else{
		for i:=0; i<len(runes); i++{
			ithIndex := runes[i]
			ros := string(runes[0:i]) + string(runes[i+1:])
			rosrune := []rune(ros)
			perm(rosrune, string(out) + string(ithIndex))
		}
	}
}