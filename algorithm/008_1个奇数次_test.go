package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddOne(t *testing.T) {
	tests := []struct {
		name string
		//输入
		arr []int

		//输出
		odd int
	}{
		{
			name: "normal case",
			arr:  []int{1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6, 7, 7},
			odd:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.odd, OddOne(tt.arr))
		})
	}
}
