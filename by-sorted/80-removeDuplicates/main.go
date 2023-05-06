package main

import "fmt"

/*原地删除有序数组中的重复元素，使得重复的元素最多出现两次*/
func removeDuplicates(nums []int) int {
	var ans int
	n := len(nums)
	if n < 2 {
		return n
	}
	for i := 0; i < n; i++ {
		l, r := i, n-1
		for l <= r {
			if nums[r] == nums[l] {
				break
			}
			r--
		}

		if r != l {
			ans += 2
		} else {
			ans += 1
		}
		i = r
		//fmt.Println(l, r, nums[l], nums[r], i, ans)

	}
	return ans
}

func main() {
	//nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}
