package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestClimbStairs -
func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "Example1",
			args: 2,
			want: 2,
		},
		{
			name: "Example2",
			args: 3,
			want: 3,
		},
		{
			name: "Example3",
			args: 45,
			want: 1836311903,
		},
	}

	for _, tt := range tests {
		res := climbStairs(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
