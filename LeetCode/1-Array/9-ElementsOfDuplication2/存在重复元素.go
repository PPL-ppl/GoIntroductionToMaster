package __ElementsOfDuplication

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := map[int]int{}
	for i, j := range nums {
		_, ok := mp[j]
		if ok {
			i2 := mp[j]
			if i-i2 <= k {
				return true
			}
		}
		mp[j] = i
	}
	return false
}
