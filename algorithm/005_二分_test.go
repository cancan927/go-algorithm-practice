package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNum(t *testing.T) {

	tests := []struct {
		name string

		// input
		arr []int
		num int
		// output
		exist bool
	}{
		{
			name:  "nil",
			arr:   nil,
			num:   9,
			exist: false,
		}, {
			name:  "empty",
			arr:   []int{},
			num:   8,
			exist: false,
		}, {
			name:  "normal true case",
			arr:   []int{1, 3, 5, 7, 9, 11, 13, 15, 17},
			num:   5,
			exist: true,
		}, {
			name:  "normal false case",
			arr:   []int{2, 4, 6, 8, 10, 12, 14, 16},
			num:   7,
			exist: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.exist, FindNum(tt.arr, tt.num))
		})
	}
}
