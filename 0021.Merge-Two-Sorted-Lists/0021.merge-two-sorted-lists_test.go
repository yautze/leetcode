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

// TestMergeTwoLists -
func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				Num1: []int{1, 2, 4},
				Num2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "Example2",
			args: args{
				Num1: []int{},
				Num2: []int{},
			},
			want: []int{},
		},
		{
			name: "Example3",
			args: args{
				Num1: []int{},
				Num2: []int{0},
			},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		res := mergeTwoLists(structures.Ints2List(tt.args.Num1), structures.Ints2List(tt.args.Num2))
		fmt.Println(res)
		assert.Equal(t, tt.want, structures.List2Ints(res))
	}
}


// TestMergeTwoLists2 -
func TestMergeTwoLists2(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				Num1: []int{1, 2, 4},
				Num2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name: "Example2",
			args: args{
				Num1: []int{},
				Num2: []int{},
			},
			want: []int{},
		},
		{
			name: "Example3",
			args: args{
				Num1: []int{},
				Num2: []int{0},
			},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		res := mergeTwoLists2(structures.Ints2List(tt.args.Num1), structures.Ints2List(tt.args.Num2))
		fmt.Println(res)
		assert.Equal(t, tt.want, structures.List2Ints(res))
	}
}
