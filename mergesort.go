package main

func MergeSort(array []int) []int {

	if len(array) <= 1 {
		return array
	} else {
		mid := len(array) / 2
		left := MergeSort(array[:mid])
		right := MergeSort(array[mid:])

		return merge(left, right)
	}

}

func merge(l, r []int) []int {
	merged := make([]int, 0, len(l)+len(r))

	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(merged, r...)
		} else if len(r) == 0 {
			return append(merged, l...)
		} else if l[0] < r[0] {
			merged = append(merged, l[0])
			l = l[1:]
		} else {
			merged = append(merged, r[0])
			r = r[1:]
		}

	}
	return merged
}
