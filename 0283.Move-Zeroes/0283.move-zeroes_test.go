package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMoveZeroes -
func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Example1",
			args: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "Example2",
			args: []int{0},
			want: []int{0},
		},
		{
			name: "Example3",
			args: []int{0, 0, 1},
			want: []int{1, 0, 0},
		},
	}

	for _, tt := range tests {
		moveZeroes(tt.args)
		fmt.Println(tt.args)
		assert.Equal(t, tt.want, tt.args)
	}
}
