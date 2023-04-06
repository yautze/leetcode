package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// args
type args struct {
	Nums   []int
	Target int
}

// TestTwoSum -
func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				Nums:   []int{2, 7, 11, 15},
				Target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "Example2",
			args: args{
				Nums:   []int{3, 2, 4},
				Target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "Example3",
			args: args{
				Nums:   []int{3, 3},
				Target: 6,
			},
			want: []int{0, 1},
		},
	}

	for _, tt := range tests {
		res := twoSum(tt.args.Nums, tt.args.Target)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
