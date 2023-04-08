package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

// TestReorderList -
func TestReorderList(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Example1",
			args: []int{1, 2, 3, 4},
			want: []int{1, 4, 2, 3},
		},
		{
			name: "Example2",
			args: []int{1, 2, 3, 4, 5},
			want: []int{1, 5, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		in := structures.Ints2List(tt.args)
		reorderList(in)
		fmt.Println(structures.List2Ints(in))
		assert.Equal(t, tt.want, structures.List2Ints(in))
	}
}
