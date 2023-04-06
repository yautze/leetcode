package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// args
type args struct {
	Nums []int
	K    int
}

// TestTopKFrequent -
func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				Nums: []int{1, 1, 1, 2, 2, 3},
				K:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "Example2",
			args: args{
				Nums: []int{1},
				K:    1,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		res := topKFrequent(tt.args.Nums, tt.args.K)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
