package main

/*
解题思路
1:双指针
2:左右求面积，左右谁小移动谁
3:判断面积是否大于目标
*/
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(height)
}
func maxArea(height []int) int {
	num := len(height)
	mianji := 0
	for left := 0; left < num; left++ {
		for right := num - 1; right > left; right-- {
			if height[right] >= height[left] {
				if height[left]*(right-left) > mianji {
					mianji = height[left] * (right - left)
				}
				break
			}
			if height[left] > height[right] {
				if height[right]*(right-left) > mianji {
					mianji = height[right] * (right - left)
				}
			}
		}
	}
	return mianji
}
