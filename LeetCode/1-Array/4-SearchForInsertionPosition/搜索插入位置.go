package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 5, 6}
	fmt.Println(searchInsert(nums, 5))
	fmt.Println(searchInsert(nums, 2))
	fmt.Println(searchInsert(nums, 7))
	fmt.Println(searchInsert(nums, 0))
	fmt.Println("-----------------")
	fmt.Println(erFen(nums, 5))
	fmt.Println(erFen(nums, 2))
	fmt.Println(erFen(nums, 7))
	fmt.Println(erFen(nums, 0))

}
func searchInsert(nums []int, target int) int {
	if nums[0] >= target {
		return 0
	}
	i2 := len(nums)
	if target > nums[i2-1] {
		return i2
	}
	index := 0
	for i, num := range nums {
		if num == target {
			index = i
			break
		}
		if i < i2 {
			if nums[i] < target && nums[i+1] > target {
				if index == 0 {
					index = i + 1
				}
			}
		} else {
			if index == 0 {
				index = i2 - 1
			}
		}

	}
	return index

}

// 二分查找
func erFen(nums []int, target int) int {
	size := len(nums)
	left := 0
	right := size - 1
	for left <= right {
		mid := left + (right-1)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
