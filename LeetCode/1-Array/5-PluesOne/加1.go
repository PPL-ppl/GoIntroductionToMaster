package main

// 解题思路
// 从后面开始遍历
// 遇到9不处理 非九当前值加1  从当前下标加1 遍历到最大长度 全部设置为0
// 全是9  新建一个长度比旧数组大1的数据，下标设置为1
func main() {
	plusOne([]int{1, 2, 3, 9})
}

func plusOne1(digits []int) []int {
	a := len(digits)
	if digits[a-1] != 9 {
		digits[a-1] = digits[a-1] + 1
		return digits
	}
	for i := a - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] == 10 {
			digits[i] = 0
			if i == 0 {
				digits[i] = 1
			}
			if i+1 == a {
				digits = append(digits, 0)
			} else {
				digits[i+1] = 0
			}
		}
	}
	return digits
}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
