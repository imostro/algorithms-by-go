package main

/*
剑指 Offer 03. 数组中重复的数字

找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。

链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
*/

// map 版本
func findRepeatNumberByMap(nums []int) int {
	var m map[int]int = make(map[int]int)

	for _, num := range nums {
		val, ok := m[num]
		if ok {
			return num
		}
		m[num] = val + 1
	}

	return -1
}

// 原地置换算法
func findRepeatNumberBySwap(nums []int) int {
	for i, num := range nums {
		if i != num && nums[num] == num {
			return num
		}
		nums[i], nums[num] = nums[num], nums[i]
	}

	return -1
}
