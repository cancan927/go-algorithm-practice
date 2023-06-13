package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRightOne(t *testing.T) {

	tests := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "zero",
			input:  0,
			output: 0,
		}, {
			name:   "127",
			input:  127,
			output: 1,
		}, {
			name:   "126",
			input:  126,
			output: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, RightOne(tt.input))
		})
	}
}
