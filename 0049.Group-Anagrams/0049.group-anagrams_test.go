package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGroupAnagramsBySort -
func TestGroupAnagramsBySort(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want [][]string
	}{
		{
			name: "Example1",
			args: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name: "Example2",
			args: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "Example3",
			args: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		res := groupAnagramsBySort(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}

// TestGroupAnagramsBySort -
func TestGroupAnagramsByCombination(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want [][]string
	}{
		{
			name: "Example1",
			args: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name: "Example2",
			args: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "Example3",
			args: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		res := groupAnagramsByCombination(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}
