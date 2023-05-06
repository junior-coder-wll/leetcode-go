//@Author: wulinlin
//@Description:
//@File:  main
//@Version: 1.0.0
//@Date: 2023/05/06 15:36

package main

import "fmt"

/* 合并两个有序数组 从后往前双指针*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	tail := m + n - 1
	for p1 >= 0 || p2 >= 0 {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
		tail--
	}
}

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	m, n := 3, 3
	merge(num1, m, num2, n)
	fmt.Println(num1)
}
