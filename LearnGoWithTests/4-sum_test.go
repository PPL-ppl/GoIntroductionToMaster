package main

import (
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	t.Run("数组形式求和", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum1(numbers)
		want := 15
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("切片求和", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum2(numbers)
		want := 15
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("输入两个数组合成一个数组", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func Sum1(numbers [5]int) int {
	sum := 0
	/*	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}*/
	//上下同理
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

func Sum2(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum2(numbers))
	}

	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum2(tail))
		}
	}

	return sums
}
