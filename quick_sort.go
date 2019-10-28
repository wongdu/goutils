package main

import "fmt"

var index int

func quickSort(s []int, left, right int) {

	if left >= right {
		return
	}

	index++
	val := s[left]
	k := left

	for idx := left + 1; idx < right; idx++ {
		if s[idx] < val {
			s[k] = s[idx]
			s[idx] = s[k+1]
			k++
		}
	}

	s[k] = val
	fmt.Println(index, " ", s, left, right)
	quickSort(s, left, k-1)
	quickSort(s, k+1, right)

}

func main() {
	// s := []int{3, 7, 4, 2, 1, 9, 6, 10}
	s := []int{3, 8, 4, 7, 1, 12, 6, 10}
	quickSort(s, 0, len(s))
	fmt.Println(s)

}
