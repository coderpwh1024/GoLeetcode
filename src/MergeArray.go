package main

import (
	"fmt"
	"sort"
)

func main() {

	var nums1 = []int{1, 2, 3, 0, 0, 0}

	var m = 3

	var nums2 = []int{2, 5, 6}

	var n = 3

	merge(nums1, m, nums2, n)

	for i := 0; i < len(nums1); i++ {
		fmt.Println(nums1[i])
	}

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i := 0; i != n; i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)

}
