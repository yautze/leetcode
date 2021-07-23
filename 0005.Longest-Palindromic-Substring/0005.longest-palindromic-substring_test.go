package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLongestPalindrome1 -
func TestLongestPalindrome1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Example1",
			args: "babad",
			want: "bab",
		},
		{
			name: "Example2",
			args: "cbbd",
			want: "bb",
		},
		{
			name: "Example3",
			args: "a",
			want: "a",
		},
		{
			name: "Example4",
			args: "ac",
			want: "a",
		},
		{
			name: "Example5",
			args: "aacabdkacaa",
			want: "aca",
		},
	}

	for _, tt := range tests {
		res := longestPalindrome1(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}

// TestLongestPalindrome2 -
func TestLongestPalindrome2(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Example1",
			args: "babad",
			want: "bab",
		},
		{
			name: "Example2",
			args: "cbbd",
			want: "bb",
		},
		{
			name: "Example3",
			args: "a",
			want: "a",
		},
		{
			name: "Example4",
			args: "ac",
			want: "a",
		},
		{
			name: "Example5",
			args: "aacabdkacaa",
			want: "aca",
		},
	}

	for _, tt := range tests {
		res := longestPalindrome2(tt.args)
		fmt.Println(res)
		assert.Equal(t, tt.want, res)
	}
}

