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

// TestRemoveElement -
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				Nums:   []int{3,2,2,3},
				Target: 3,
			},
			want: 2,
		},
		{
			name: "Example2",
			args: args{
				Nums:   []int{0,1,2,2,3,0,4,2},
				Target: 2,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		res := removeElement(tt.args.Nums, tt.args.Target)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
