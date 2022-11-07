package main

func main() {
	nums := []int{0, 1, 0}
	canPlaceFlowers1(nums, 1)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	m := len(flowerbed)
	prev := -1
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			if prev < 0 {
				count += i / 2
			} else {
				count += (i - prev - 2) / 2
			}
			prev = i
		}
	}
	if prev < 0 {
		count += (m + 1) / 2
	} else {
		count += (m - prev - 1) / 2
	}
	return count >= n
}
func canPlaceFlowers1(flowerbed []int, n int) bool {
	num := len(flowerbed)
	if num == 1 {
		if flowerbed[0] == 0 {
			n--
		}
		return n <= 0
	}
	if num == 2 {
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			n--
		}
		return n <= 0
	}
	for i := 0; i < num; i++ {
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			flowerbed[0] = 1
			i = 1
			n--
			continue
		}
		if i == num-1 && flowerbed[num-2] == 0 && flowerbed[num-1] == 0 {
			flowerbed[num-1] = 1
			n--
			break
		}
		if i == num-1 {
			break
		}
		if i != 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
			flowerbed[i] = 1
			n--
			continue
		}
	}
	return n <= 0
}
