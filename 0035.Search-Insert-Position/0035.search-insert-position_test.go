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

// TestSearchInsert -
func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				Nums:   []int{1, 3, 5, 6},
				Target: 5,
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				Nums:   []int{1, 3, 5, 6},
				Target: 2,
			},
			want: 1,
		},
		{
			name: "Example3",
			args: args{
				Nums:   []int{1, 3, 5, 6},
				Target: 7,
			},
			want: 4,
		},
		{
			name: "Example4",
			args: args{
				Nums:   []int{1, 3, 5, 6},
				Target: 0,
			},
			want: 0,
		},
		{
			name: "Example3",
			args: args{
				Nums:   []int{1},
				Target: 0,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		res := searchInsert(tt.args.Nums, tt.args.Target)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
