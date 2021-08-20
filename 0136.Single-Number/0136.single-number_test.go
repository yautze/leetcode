package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSingleNumberByMap -
func TestSingleNumberByMap(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example1",
			args: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "Example2",
			args: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "Example3",
			args: []int{1},
			want: 1,
		},
	}

	for _, tt := range tests {
		res := singleNumberByMap(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}

// TestSingleNumber -
func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example1",
			args: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "Example2",
			args: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "Example3",
			args: []int{1},
			want: 1,
		},
	}

	for _, tt := range tests {
		res := singleNumber(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}