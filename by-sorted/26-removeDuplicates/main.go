package main

/*
1.比较 p 和 q 位置的元素是否相等。

如果相等，q 后移 1 位 如果不相等，将 q 位置的元素复制到 p+1 位置上，p 后移一位，q 后移 1 位 重复上述过程，直到 q 等于数组长度。

返回 p + 1，即为新数组长度。
*/

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	p := 0
	q := p + 1
	for q < n {
		if nums[p] != nums[q] {
			nums[p+1] = nums[q]
			p++
		}
		q++
	}
	return p + 1
}

func main() {

}
