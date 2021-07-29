package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
	"github.com/stretchr/testify/assert"
)

// args
type args struct {
	Num1 []int
	N    int
}

// TestRemoveNthFromEnd -
func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				Num1: []int{1, 2, 3, 4, 5},
				N:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "Example2",
			args: args{
				Num1: []int{1},
				N:    1,
			},
			want: []int{},
		},
		{
			name: "Example3",
			args: args{
				Num1: []int{1, 2},
				N:    1,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		res := removeNthFromEnd(structures.Ints2List(tt.args.Num1), tt.args.N)
		fmt.Println(res)
		assert.Equal(t, tt.want, structures.List2Ints(res))
	}
}
