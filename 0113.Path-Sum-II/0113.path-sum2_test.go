package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

type args struct {
	Num    []int
	Target int
}

// TestPathSum -
func TestPathSum(t *testing.T) {
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example1",
			args: args{
				Num:    []int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, 5, 1},
				Target: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name: "Example2",
			args: args{
				Num:    []int{1, 2, 3},
				Target: 5,
			},
			want: [][]int{},
		},
		{
			name: "Example3",
			args: args{
				Num:    []int{1, 2},
				Target: 0,
			},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		res := pathSum(structures.Ints2TreeNode(tt.args.Num), tt.args.Target)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
