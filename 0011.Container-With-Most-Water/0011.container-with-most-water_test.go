package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMaxArea -
func TestMaxArea(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example1",
			args: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want: 49,
		},
		{
			name: "Example2",
			args: []int{1, 1},
			want: 1,
		},
		{
			name: "Example3",
			args: []int{4, 3, 2, 1, 4},
			want: 16,
		},
		{
			name: "Example3",
			args: []int{1, 2, 1},
			want: 2,
		},
	}

	for _, tt := range tests {
		res := maxArea(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
