package main

import (
	"testing"
)

func Test_Add(t *testing.T) {
	//assert.Equal(t, 4, actual) //expected 预期  actual实际值
	t.Run("加法测试", func(t *testing.T) {
		actual := Add(2, 2)
		expected := 4
		if expected != actual {
			t.Errorf("expected '%d' but got '%d'", expected, actual)
		}
	})
}
func Add(a, b int) int {
	return a + b
}
