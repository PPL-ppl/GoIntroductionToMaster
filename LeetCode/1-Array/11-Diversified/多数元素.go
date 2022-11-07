package main

func main() {
	num := []int{8, 8, 7, 7, 7}
	majorityElement(num)
}
func majorityElement(nums []int) int {
	index := 0
	ru := -1
	mp := map[int]map[int]int{}
	for _, i2 := range nums {
		_, i3 := mp[i2][i2]
		if i3 {
			mp[i2][i2] = mp[i2][i2] + 1
			if mp[i2][i2] >= index {
				index = mp[i2][i2]
				ru = i2
			}
		} else {
			m := map[int]int{}
			m[i2] = 0
			mp[i2] = m
			if ru == -1 {
				ru = i2
			}
		}
	}
	return ru
}

//众数原理

func majorityElement1(nums []int) int {
	candidate := 0
	count := 0 //出现次数
	for _, i2 := range nums {
		if count == 0 {
			candidate = i2
		}
		if i2 == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
