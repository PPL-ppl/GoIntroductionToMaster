package main

import "fmt"

func main() {
	nums := []int{4, 2, 1, 3, 2, 4, 3}
	number := singleNumber1(nums)
	fmt.Println(number)
}

func singleNumber(nums []int) int {
	index := 0

	number := make([]int, len(nums))
	ins := make(map[int]int)
	for _, v := range nums {
		_, ok := ins[v]
		if ok {
			ins[v] = ins[v] + 1

		} else {
			number = append(number, v)
			ins[v] = 1
		}
	}
	for i := range ins {
		if ins[i] == 1 {
			index = i
		}
	}
	return index
}

// 异或运算
func singleNumber1(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
