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
func TestSearchByMap(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				Nums:   []int{4, 5, 6, 7, 0, 1, 2},
				Target: 0,
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				Nums:   []int{4, 5, 6, 7, 0, 1, 2},
				Target: 3,
			},
			want: -1,
		},
		{
			name: "Example3",
			args: args{
				Nums:   []int{1},
				Target: 0,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		res := searchByMap(tt.args.Nums, tt.args.Target)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}

// TestSearch -
func TestSearch(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				Nums:   []int{4, 5, 6, 7, 0, 1, 2},
				Target: 0,
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				Nums:   []int{4, 5, 6, 7, 0, 1, 2},
				Target: 3,
			},
			want: -1,
		},
		{
			name: "Example3",
			args: args{
				Nums:   []int{1},
				Target: 0,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		res := search(tt.args.Nums, tt.args.Target)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}