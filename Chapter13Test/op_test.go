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
