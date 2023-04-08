package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMissingNumberByXOR -
func TestMissingNumberByXOR(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example1",
			args: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "Example2",
			args: []int{0, 1},
			want: 2,
		},
		{
			name: "Example3",
			args: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
	}

	for _, tt := range tests {
		res := missingNumberByXOR(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}

// TestMissingNumberByCyclicSort -
func TestMissingNumberByCyclicSort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Example1",
			args: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "Example2",
			args: []int{0, 1},
			want: 2,
		},
		{
			name: "Example3",
			args: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
	}

	for _, tt := range tests {
		res := missingNumberByCyclicSort(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}