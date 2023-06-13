package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumKLength(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want int
	}{
		{
			name: "nil case",
			arr:  nil,
			k:    9,
			want: -1,
		}, {
			name: "nil 0 case",
			arr:  nil,
			k:    0,
			want: -1,
		}, {
			name: "empty k case",
			arr:  []int{},
			k:    9,
			want: -1,
		}, {
			name: "normal case",
			arr:  []int{1, 2, 3, 4, 5, 6},
			k:    9,
			want: 3,
		}, {
			name: "0 case",
			arr:  []int{5, 1, 4, 3},
			k:    0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SumKLength(tt.arr, tt.k))
		})
	}
}
