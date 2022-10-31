package main

import "fmt"

/*
给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

从0开始遍历 记下不等于目标数的值后面一个下标 遇到相等跳过  后面遇到放到该位置
*/
func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(easty(nums, 2))
}

func removeElement(nums []int, val int) int {
	a := 0
	i := len(nums)
	for slow := 0; slow < i; slow++ {
		if nums[slow] != val {
			a++
		}
		for fast := i - 1; fast > slow; fast-- {
			if nums[slow] == val {
				if nums[fast] != val {
					nums[slow], nums[fast] = nums[fast], nums[slow]
					a++
				}
			}
		}
	}

	return a
}

func easty(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
