package __ElementsOfDuplication

func containsDuplicate(nums []int) bool {
	mp := map[int]int{}
	for _, j := range nums {
		_, ok := mp[j]
		if ok {
			return true
		}
	}
	return false
}
