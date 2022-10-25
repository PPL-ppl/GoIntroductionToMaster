package main

//需求 重复打印机传入的值
import "testing"

func Test_for(t *testing.T) {

	t.Run("测试For循环", func(t *testing.T) {
		actula := Repeat("a")
		expected := "aaaaa"
		if actula != expected {
			t.Errorf("expected:%q,expeactde:%q", actula, expected)
		}
	})
}

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

// 基准测试
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
