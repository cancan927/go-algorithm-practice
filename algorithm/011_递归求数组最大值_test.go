package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getMax(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		max  int
	}{
		{
			name: "one element",
			arr:  []int{1},
			max:  1,
		}, {
			name: "normal case",
			arr:  []int{1, 2, 3, 4, 5},
			max:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.max, getMax(tt.arr))
		})
	}
}
