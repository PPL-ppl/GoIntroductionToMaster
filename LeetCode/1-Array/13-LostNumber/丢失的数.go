package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	nums := []int{3, 2, 1}
	missingNumber(nums)
	ss()
}
func missingNumber(nums []int) int {
	data := []int{10, 22, 11, 24, 17, 26}
	index := sort.Search(len(data), func(i int) bool { return data[i] > 23 })
	fmt.Println(index, data[index]) // 3 24
	for index, v := range nums {
		if index != v {
			return index
		}
	}
	return len(nums)
}
func ss() {
	x := []string{"1", "a", "1", "a"}
	target := "a"
	//使用find的前提是已经排好序
	sort.Strings(x)
	i, found := sort.Find(len(x), func(i int) int {
		return strings.Compare(target, x[i])
	})
	if found {
		fmt.Printf("found %s at entry %d\n", target, i)
	} else {
		fmt.Printf("%s not found, would insert at %d", target, i)
	}
}
