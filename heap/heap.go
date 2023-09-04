package heap

import "fmt"

type heap struct {
	list []int
}

func New() *heap {
	return &heap{list: make([]int, 0)}
}

func (h *heap) Add(val int) {
	h.list = append(h.list, val)
	h.heapifyUp()
}

func (h *heap) Extract() int {
	if len(h.list) == 0 {
		return -1
	}
	value := h.list[0]
	h.list[0] = h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	h.heapifyDown()
	return value
}

func (h *heap) Print() {
	fmt.Printf("%v\n", h.list)
}

func (h *heap) heapifyUp() {

	l := h.list
	if len(l) < 2 {
		return
	}
	index := len(l) - 1
	root := parent(index)

	for l[index] > l[root] {
		l[index], l[root] = l[root], l[index]
		index = root
		root = parent(index)
	}

}

func (h *heap) heapifyDown() {
	list := h.list
	root := 0
	l := left(root)
	r := right(root)

	for {

		if l >= len(list) || r >= len(list) {
			break
		}
		if l == len(list)-1 {
			if list[root] < list[l] {
				list[root], list[l] = list[l], list[root]
				break
			}
		}
		if list[l] > list[r] {
			if list[root] < list[l] {
				list[root], list[l] = list[l], list[root]
				root = l
				l = left(root)
				r = right(root)
				continue
			}
		} else {
			if list[root] < list[r] {
				list[root], list[r] = list[r], list[root]
				root = r
				l = left(root)
				r = right(root)
				continue
			}
		}
		break
	}
}

func left(root int) int {
	return (2 * root) + 1
}

func right(root int) int {
	return (2 * root) + 2
}

func parent(index int) int {
	return (index - 1) / 2
}
