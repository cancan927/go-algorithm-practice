package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {

	tests := []struct {
		name string
		in   []int
		p    int
		want []int
	}{
		{
			name: "nil",
			in:   nil,
			p:    1,
			want: nil,
		}, {
			name: "1 element",
			in:   []int{3},
			p:    2,
			want: []int{3},
		}, { // 这个测试不严谨
			name: "normal",
			in:   []int{7, 3, 2, 9, 5, 6, 10},
			p:    6,
			want: []int{3, 2, 5, 6, 7, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Partition(tt.in, tt.p))
		})
	}
}
