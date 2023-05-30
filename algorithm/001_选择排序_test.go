package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name string

		//输入
		input []int

		//输出
		want []int
	}{
		{
			name:  "nil",
			input: nil,
			want:  nil,
		}, {
			name:  "len<2",
			input: []int{2},
			want:  []int{2},
		}, {
			name:  "normal case",
			input: []int{3, 1, 5, 2, 9, 6, 0, 5},
			want:  []int{0, 1, 2, 3, 5, 5, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.input)
			assert.Equal(t, tt.input, tt.want)
		})
	}
}
