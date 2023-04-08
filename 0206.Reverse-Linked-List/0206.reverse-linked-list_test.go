package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

// TestReverseList -
func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Example1",
			args: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "Example2",
			args: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "Example3",
			args: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		res := reverseList(structures.Ints2List(tt.args))
		fmt.Println(res)
		assert.Equal(t, tt.want, structures.List2Ints(res))
	}
}
