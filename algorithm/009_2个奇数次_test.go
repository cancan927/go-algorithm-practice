package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddTwo(t *testing.T) {

	tests := []struct {
		name  string
		arr   []int
		wantA int
		wantB int
	}{
		{
			name:  "case 1",
			arr:   []int{1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 5, 5, 6, 6, 9, 11, 7, 7},
			wantA: 11,
			wantB: 9,
		}, {
			name:  "case 2",
			arr:   []int{1, 4, 3, 1, 7, 2, 7, 2, 3, 1, 2, 7, 2, 3, 1, 3, 4, 1},
			wantA: 7,
			wantB: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := OddTwo(tt.arr)
			assert.Equal(t, tt.wantA, gotA)
			assert.Equal(t, tt.wantB, gotB)
		})
	}
}
