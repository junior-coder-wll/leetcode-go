package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	arrLen := len(nums)
	var ans []int
	// 1. 特殊值判断
	if arrLen < 2 {
		return ans
	}
	// 2. 创建哈希表，用于保存数组中的所有元素，Key为数组元素的值，value为数组元素的值下标
	hashMap := make(map[int]int, arrLen)
	// 3. 遍历数组，将数据填充哈希表
	for i := 0; i < arrLen; i++ {
		hashMap[nums[i]] = i
	}
	// 4. 遍历哈希表，和target进行运算，求出对应结果
	for i := 0; i < arrLen; i++ {
		res := target - nums[i]
		if v, ok := hashMap[res]; ok && v != i {
			ans = append(ans, i)
			ans = append(ans, v)
			return ans
		}
	}
	return ans
}
