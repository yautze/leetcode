package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example1",
			args: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "Example2",
			args: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "Example3",
			args: []int{7, 8, 9, 11, 12},
			want: 1,
		},
	}

	for _, tt := range tests {
		res := firstMissingPositive(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
