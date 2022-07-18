package Chapter13Test

import "testing"

func sum() int {
    total := 0
    for i := 0; i < 10; i++ {
        total += i
    }
    return total
}
func BenchmarkSum(b testing.B) {
    for i := 0; i < 100; i++ {
        sum()
    }
}
