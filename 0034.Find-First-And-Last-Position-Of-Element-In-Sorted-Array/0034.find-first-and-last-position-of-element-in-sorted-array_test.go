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

// TestSearchByMap -
func TestSearchRanges(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				Nums:   []int{5,7,7,8,8,10},
				Target: 8,
			},
			want: []int{3,4},
		},
		{
			name: "Example2",
			args: args{
				Nums:   []int{5,7,7,8,8,10},
				Target: 6,
			},
			want: []int{-1,-1},
		},
		{
			name: "Example3",
			args: args{
				Nums:   []int{},
				Target: 0,
			},
			want: []int{-1,-1},
		},
	}

	for _, tt := range tests {
		res := searchRange(tt.args.Nums, tt.args.Target)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}