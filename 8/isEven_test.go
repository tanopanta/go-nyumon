package evenodd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEven(t *testing.T) {
	tc := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "偶数",
			num:      2,
			expected: true,
		},
		{
			name:     "奇数",
			num:      1,
			expected: false,
		},
		{
			name:     "0",
			num:      0,
			expected: true,
		},
		{
			name:     "マイナスの場合",
			num:      -1,
			expected: false,
		},
	}

	for _, c := range tc {
		c := c

		t.Run(c.name, func(t *testing.T) {
			//事前準備
			actual := isEven(c.num)
			assert.Equal(t, c.expected, actual)
		})
	}
}
