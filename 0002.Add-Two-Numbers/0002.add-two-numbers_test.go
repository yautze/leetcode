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
	Num2 []int
}

func Test_Problem2(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				Num1: []int{2, 4, 3},
				Num2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "Example2",
			args: args{
				Num1: []int{0},
				Num2: []int{0},
			},
			want: []int{0},
		},
		{
			name: "Example3",
			args: args{
				Num1: []int{9, 9, 9, 9, 9, 9, 9},
				Num2: []int{9, 9, 9, 9},
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		res := addTwoNumbers(structures.Ints2List(tt.args.Num1), structures.Ints2List(tt.args.Num2))
		fmt.Println(res)
		assert.Equal(t, tt.want, structures.List2Ints(res))
	}
}
