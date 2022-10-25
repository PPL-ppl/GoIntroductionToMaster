package Chapter13Test

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := sum()
	expected := 4
	//assert.Equal(t, expected, actual)
	if expected != actual {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}
func sum() int {
	total := 0
	for i := 0; i < 10; i++ {
		total += i
	}
	return total
}
