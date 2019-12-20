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

	fmt.Println("the to sort is:", s[left:right])
	for idx := left + 1; idx < right; idx++ {
		if s[idx] < val {
			fmt.Println("idx=", idx, "val=", val, "k=", k)
			s[k] = s[idx]
			s[idx] = s[k+1]
			k++
		}
		fmt.Println("after one compare:", s)
	}

	s[k] = val
	fmt.Println(index, " ", s, left, right)
	quickSort(s, left, k-1)
	quickSort(s, k+1, right)

}

func main() {
	// s := []int{3, 7, 4, 2, 1, 9, 6, 10}
	s := []int{3, 1, 4, 7, 2, 12, 6, 10}

	quickSort(s, 0, len(s))
	fmt.Println(s)

}
