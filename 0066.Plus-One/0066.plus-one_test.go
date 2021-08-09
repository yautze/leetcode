package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPlusOne -
func TestPlusOne(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Example1",
			args: []int{1, 2, 3},
			want: []int{1, 2, 4},
		},
		{
			name: "Example2",
			args: []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "Example3",
			args: []int{0},
			want: []int{1},
		},
		{
			name: "Example4",
			args: []int{9, 9, 9, 9},
			want: []int{1, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		res := plusOne(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}

}
