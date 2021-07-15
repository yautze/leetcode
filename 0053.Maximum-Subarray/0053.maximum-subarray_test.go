package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMaxSubArray -
func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example1",
			args: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "Example2",
			args: []int{1},
			want: 1,
		},
		{
			name: "Example3",
			args: []int{5, 4, -1, 7, 8},
			want: 23,
		},
	}

	for _, tt := range tests {
		res := maxSubArray(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
