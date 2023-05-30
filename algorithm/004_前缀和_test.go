package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefixSum(t *testing.T) {
	tests := []struct {
		name string

		// input
		arr []int
		l   int
		r   int

		//output
		sum int
	}{
		{
			name: "one elements",
			arr:  []int{6},
			l:    0,
			r:    0,
			sum:  6,
		},
		{
			name: "normal case",
			arr:  []int{1, 2, 3, 4, 5, 6},
			l:    1,
			r:    4,
			sum:  14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.sum, PrefixSum(tt.arr, tt.l, tt.r))
		})
	}
}
