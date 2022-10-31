package main

import (
	"fmt"
	"time"
)

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
示例1
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例2
输入：nums = [3,2,4], target = 6
输出：[1,2]
示例3
输入：nums = [3,3], target = 6
输出：[0,1]
*/
func main() {
	m := []int{2, 3, 4, 7, 8, 9, 0, 66, 55, 444, 333, 2222, 111111}
	fmt.Println(time.Now().String())
	sumTarget(m, 113333)
	fmt.Println(time.Now().String())

	fmt.Println("--------------")
	fmt.Println(time.Now().String())
	sumTargetHashSet(m, 113333)
	fmt.Println(time.Now().String())

}
func sumTarget(nums []int, target int) []int {
	m := []int{}
	i2 := len(nums)

	for i := 0; i <= i2; i++ {
		for y := i + 1; y < i2; y++ {
			if (nums[y] + nums[i]) == target {
				m = append(m, y, i)
				return m
			}
		}
	}
	return nil
}

func sumTargetHashSet(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
