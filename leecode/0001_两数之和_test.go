package leecode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {

	tests := []struct {
		name   string
		input1 []int
		input2 int
		want   []int
	}{
		{
			name:   "nil case",
			input1: []int{2, 3, 5, 8},
			input2: 29,
			want:   nil,
		}, {
			name:   "normal case",
			input1: []int{2, 7, 11, 15},
			input2: 9,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := twoSum(tt.input1, tt.input2)
			assert.Equal(t, tt.want, res)
		})
	}
}
