package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestContainsDuplicate -
func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example1",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "Example2",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "Example3",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
	}

	for _, tt := range tests {
		res := containsDuplicate(tt.nums)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
