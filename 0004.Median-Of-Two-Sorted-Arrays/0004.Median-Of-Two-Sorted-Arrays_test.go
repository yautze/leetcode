package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// args
type args struct {
	Nums1 []int
	Nums2 []int
}

// TestFindMedianSortedArrays -
func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example1",
			args: args{
				Nums1: []int{1, 3},
				Nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				Nums1: []int{1, 2},
				Nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			name: "Example3",
			args: args{
				Nums1: []int{0,0},
				Nums2: []int{0,0},
			},
			want: 0,
		},
		{
			name: "Example4",
			args: args{
				Nums1: []int{},
				Nums2: []int{1},
			},
			want: 1,
		},
		{
			name: "Example5",
			args: args{
				Nums1: []int{2},
				Nums2: []int{},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		res := findMedianSortedArrays(tt.args.Nums1, tt.args.Nums2)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
