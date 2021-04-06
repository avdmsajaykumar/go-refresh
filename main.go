package main

import "fmt"

type sum struct {
	a int32
	b int32
}

func (s *sum) add() int32 {
	return s.a + s.b
}

func main() {

	// z := sum{a: 10, b: 20}
	// fmt.Println(z.add())

	strin := "Hello World, How are you, I'm doing great"

	fmt.Println(SentenceReverse(strin))
	fmt.Println(StringReverse(strin))

	data := [] int{9,4,3,6,1,2,10,5,7,8}
	fmt.Printf("%v\n%v\n", data, MergeSort(data))

	Permutation("abc")

}