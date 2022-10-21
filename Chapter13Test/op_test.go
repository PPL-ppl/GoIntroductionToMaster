package Chapter13Test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	result := sum()
	fmt.Println(result)
	assert.Equal(t, result, 45)
}
func sum() int {
	total := 0
	for i := 0; i < 10; i++ {
		total += i
	}
	return total
}
