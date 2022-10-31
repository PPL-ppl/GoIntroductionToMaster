package main

//双指针
import "sort"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, len(nums1), nums2, len(nums2))
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		//当指针1到了m 但是n还没有结束  就把nums2的追加到后面
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		//当指针2到了n 但是m还没有结束  就把nums1剩下的追加到后面
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			//移动指针
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			//移动指针
			p2++
		}
	}
	copy(nums1, sorted)
}
