package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCanJump -
func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want bool
	}{
		{
			name: "Example1",
			args: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "Example2",
			args: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "Example3",
			args: []int{1},
			want: true,
		},
	}

	for _, tt := range tests {
		res := canJump(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
