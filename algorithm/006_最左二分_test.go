package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMostLeft(t *testing.T) {
	tests := []struct {
		name string

		//input
		arr []int
		num int
		//output
		want int
	}{
		{
			name: "nil",
			arr:  nil,
			num:  3,
			want: -1,
		}, {
			name: "empty",
			arr:  []int{},
			num:  3,
			want: -1,
		}, {
			name: "normal case",
			arr:  []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
			num:  8,
			want: 4,
		}, {
			name: "another case",
			arr:  []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			num:  8,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MostLeft(tt.arr, tt.num))
		})
	}
}
